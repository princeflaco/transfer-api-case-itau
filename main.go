package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	"transfer-api/adapter/controller"
	"transfer-api/adapter/repository"
	repository2 "transfer-api/core/repository"
	"transfer-api/core/usecase"
	"transfer-api/infra/server"
)

const (
	Port     = 8080
	Timeout  = 30 * time.Second
	BasePath = "/api/v1"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	ginServer := server.NewGinServer(Port, Timeout).SetBasePath(BasePath)

	// adicionando rotas
	ginServer.AddHandler(http.MethodGet, "/transfers/:accountId", getTransferHistoryHandler)
	ginServer.AddHandler(http.MethodPost, "/transfers/:accountId", createTransferHandler)
	ginServer.AddHandler(http.MethodGet, "/customers", getAllCustomersHandler)
	ginServer.AddHandler(http.MethodGet, "/customers/:accountId", getCustomerHandler)
	ginServer.AddHandler(http.MethodPost, "/customers", createCustomerHandler)

	ginServer.ListenAndServe(ctx, &wg)
	wg.Wait()
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
