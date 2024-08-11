package usecase

import (
	"context"
	"go.uber.org/zap"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type GetCustomerUseCase struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
}

func NewGetCustomerUseCase(customerRepo repository.CustomerRepository, accountRepo repository.AccountRepository) *GetCustomerUseCase {
	return &GetCustomerUseCase{
		CustomerRepo: customerRepo,
		AccountRepo:  accountRepo,
	}
}

func (uc *GetCustomerUseCase) Execute(ctx context.Context, accountId string) (*output.GetCustomerOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start get customer use case")
	defer log.Info("end get customer use case")

	account, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		log.Error("error while retrieving account from database", zap.Error(err))
		return nil, err
	}

	customer, err := uc.CustomerRepo.GetById(account.CustomerId)
	if err != nil {
		log.Error("error while retrieving customer from database", zap.Error(err))
		return nil, err
	}
	log.Debug("retrieved customer from database", zap.Any("customer", customer))

	balance := util.CentsToFloat64(account.Balance)
	customerOutput := output.NewGetCustomerOutput(customer.Id, customer.Name, account.Id, balance)
	return customerOutput, nil
}
