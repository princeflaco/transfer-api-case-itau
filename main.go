package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sync"
	"time"
	"transfer-api/adapter/controller"
	repoImpl "transfer-api/adapter/repository"
	repo "transfer-api/core/repository"
	"transfer-api/core/service"
	"transfer-api/core/usecase"
	"transfer-api/docs"
	"transfer-api/infra"
	"transfer-api/infra/server"
)

const (
	Version  = "v1"
	BasePath = "/api/" + Version
)

//	@contact.name	Desenvolvedor
//	@contact.url	https://github.com/princeflaco
//	@contact.email	leobbispo@hotmail.com

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// iniciando contexto
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// carregando variáveis de ambiente
	err := infra.LoadConfig()
	if err != nil {
		panic(errors.New("error while initializing config: " + err.Error()))
	}

	// iniciando logger
	logger := infra.NewLogger(infra.Config.LoggingLevel)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Sprintf("sync logger failed: %v", err))
		}
	}(logger)
	logger.Info("logger initialized")
	logger.Info("config loaded", zap.Any("config", infra.Config))

	// adicionando logger ao contexto
	ctx = context.WithValue(ctx, "logger", logger)

	// configurando valores estáticos do swagger
	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.BasePath = BasePath
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = "localhost:" + infra.Config.Port
	docs.SwaggerInfo.Title = infra.Config.AppName
	docs.SwaggerInfo.Description = "Api de transferências"

	var wg sync.WaitGroup

	// pegando variáveis de porta e timeout da infra
	port, err := strconv.Atoi(infra.Config.Port)
	if err != nil {
		panic(fmt.Errorf("failed to convert port to int: %v", err))
	}
	logger.Info("port selected", zap.Int("port", port))

	timeout := time.Duration(infra.Config.Timeout) * time.Second
	logger.Info("timeout selected", zap.Duration("timeout", timeout))

	// criando configurações da transferência
	transferConfig := usecase.TransferConfig{
		MaxAmount:   float64(infra.Config.TransferMaxAmount),
		WorkerCount: infra.Config.TransferWorkerCount,
	}
	logger.Info("transfer config set", zap.Any("config", transferConfig))

	// iniciando repositórios
	customerRepoImpl := repoImpl.NewInMemCustomerRepository()
	accountRepoImpl := repoImpl.NewInMemAccountRepository()
	transferRepoImpl := repoImpl.NewInMemTransferRepository()
	logger.Info("all repository initialized")

	// iniciando serviço
	transferServiceImpl := service.NewTransferServiceImpl(transferRepoImpl, accountRepoImpl)
	logger.Info("all service initialized")

	// criando handlers
	createCustomerHandler := setupCreateCustomerHandler(customerRepoImpl, accountRepoImpl)
	getCustomerHandler := setupGetCustomerHandler(customerRepoImpl, accountRepoImpl)
	getAllCustomersHandler := setupGetAllCustomersHandler(customerRepoImpl, accountRepoImpl)
	createTransferHandler := setupCreateTransferHandler(transferServiceImpl, transferConfig)
	getTransferHistoryHandler := setupGetTransferHistoryHandler(transferRepoImpl)
	logger.Info("all handler initialized")

	// iniciando servidor
	ginServer := server.NewGinServer(int64(port), timeout).SetBasePath(BasePath)
	logger.Info("gin server initialized")

	// adicionando middleware para o contexto e request-id
	ginServer.Engine.Use(ContextMiddleware(logger))
	logger.Info("gin middleware set")

	// adicionando rotas
	ginServer.AddHandler(http.MethodGet, "/transfers/:accountId", getTransferHistoryHandler)
	ginServer.AddHandler(http.MethodPost, "/transfers/:accountId", createTransferHandler)
	ginServer.AddHandler(http.MethodGet, "/customers", getAllCustomersHandler)
	ginServer.AddHandler(http.MethodGet, "/customers/:accountId", getCustomerHandler)
	ginServer.AddHandler(http.MethodPost, "/customers", createCustomerHandler)
	logger.Info("gin server handlers set")

	// escutando requisições
	ginServer.ListenAndServe(ctx, &wg)
	logger.Info("gin server listening started...")

	wg.Wait()
}

func ContextMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "logger", logger)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func setupCreateCustomerHandler(customerRepo repo.CustomerRepository, accountRepo repo.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewCreateCustomerUseCase(customerRepo, accountRepo)
	cn := controller.NewCreateCustomerController(*uc)
	return func(c *gin.Context) {
		cn.Execute(c.Writer, c.Request)
	}
}

func setupGetCustomerHandler(customerRepo repo.CustomerRepository, accountRepo repo.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewGetCustomerUseCase(customerRepo, accountRepo)
	cn := controller.NewGetCustomerController(*uc)
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		query := c.Request.URL.Query()
		query.Add("accountId", accountId)
		c.Request.URL.RawQuery = query.Encode()
		cn.Execute(c.Writer, c.Request)
	}
}

func setupCreateTransferHandler(transferService service.TransferService, config usecase.TransferConfig) gin.HandlerFunc {
	uc := usecase.NewCreateTransferUseCase(transferService, config)
	defer uc.Shutdown()
	cn := controller.NewCreateTransferController(uc)
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		query := c.Request.URL.Query()
		query.Add("accountId", accountId)
		c.Request.URL.RawQuery = query.Encode()
		cn.Execute(c.Writer, c.Request)
	}
}

func setupGetTransferHistoryHandler(transferRepo repo.TransferRepository) gin.HandlerFunc {
	uc := usecase.NewGetTransferHistoryUseCase(transferRepo)
	cn := controller.NewGetTransferHistoryController(*uc)
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		query := c.Request.URL.Query()
		query.Add("accountId", accountId)
		c.Request.URL.RawQuery = query.Encode()
		cn.Execute(c.Writer, c.Request)
	}
}

func setupGetAllCustomersHandler(customerRepo repo.CustomerRepository, accountRepo repo.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewGetAllCustomersUseCase(customerRepo, accountRepo)
	cn := controller.NewGetAllCustomersController(*uc)
	return func(c *gin.Context) {
		cn.Execute(c.Writer, c.Request)
	}
}
