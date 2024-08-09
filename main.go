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
	"transfer-api/adapter/repository"
	repository2 "transfer-api/core/repository"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := infra.LoadConfig()
	if err != nil {
		panic(errors.New("error while initializing config: " + err.Error()))
	}

	logger := infra.NewLogger(infra.Config.LoggingLevel)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Sprintf("sync logger failed: %v", err))
		}
	}(logger)

	ctx = context.WithValue(ctx, "logger", logger)

	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.BasePath = BasePath
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = "localhost:" + infra.Config.Port
	docs.SwaggerInfo.Title = infra.Config.AppName
	docs.SwaggerInfo.Description = "Api de transferências"

	var wg sync.WaitGroup

	// iniciando repositórios
	customerRepoImpl := repository.NewInMemCustomerRepository()
	accountRepoImpl := repository.NewInMemAccountRepository()
	transferRepoImpl := repository.NewInMemTransferRepository()

	// criando handlers
	createCustomerHandler := setupCreateCustomerHandler(customerRepoImpl, accountRepoImpl)
	getCustomerHandler := setupGetCustomerHandler(customerRepoImpl, accountRepoImpl)
	getAllCustomersHandler := setupGetAllCustomersHandler(customerRepoImpl, accountRepoImpl)
	createTransferHandler := setupCreateTransferHandler(transferRepoImpl, accountRepoImpl)
	getTransferHistoryHandler := setupGetTransferHistoryHandler(transferRepoImpl)

	port, err := strconv.Atoi(infra.Config.Port)
	if err != nil {
		panic(fmt.Errorf("failed to convert port to int: %v", err))
	}
	timeout := time.Duration(infra.Config.Timeout) * time.Second
	ginServer := server.NewGinServer(int64(port), timeout).SetBasePath(BasePath)

	ginServer.Engine.Use(ContextMiddleware(logger))

	// adicionando rotas
	ginServer.AddHandler(http.MethodGet, "/transfers/:accountId", getTransferHistoryHandler)
	ginServer.AddHandler(http.MethodPost, "/transfers/:accountId", createTransferHandler)
	ginServer.AddHandler(http.MethodGet, "/customers", getAllCustomersHandler)
	ginServer.AddHandler(http.MethodGet, "/customers/:accountId", getCustomerHandler)
	ginServer.AddHandler(http.MethodPost, "/customers", createCustomerHandler)

	ginServer.ListenAndServe(ctx, &wg)
	wg.Wait()
}

func ContextMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "logger", logger)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func setupCreateCustomerHandler(customerRepo repository2.CustomerRepository, accountRepo repository2.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewCreateCustomerUseCase(customerRepo, accountRepo)
	cn := controller.NewCreateCustomerController(*uc)
	return func(c *gin.Context) {
		cn.Execute(c.Writer, c.Request)
	}
}

func setupGetCustomerHandler(customerRepo repository2.CustomerRepository, accountRepo repository2.AccountRepository) gin.HandlerFunc {
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

func setupCreateTransferHandler(transferRepo repository2.TransferRepository, accountRepo repository2.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewCreateTransferUseCase(transferRepo, accountRepo)
	cn := controller.NewCreateTransferController(uc)
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		query := c.Request.URL.Query()
		query.Add("accountId", accountId)
		c.Request.URL.RawQuery = query.Encode()
		cn.Execute(c.Writer, c.Request)
	}
}

func setupGetTransferHistoryHandler(transferRepo repository2.TransferRepository) gin.HandlerFunc {
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

func setupGetAllCustomersHandler(customerRepo repository2.CustomerRepository, accountRepo repository2.AccountRepository) gin.HandlerFunc {
	uc := usecase.NewGetAllCustomersUseCase(customerRepo, accountRepo)
	cn := controller.NewGetAllCustomersController(*uc)
	return func(c *gin.Context) {
		cn.Execute(c.Writer, c.Request)
	}
}
