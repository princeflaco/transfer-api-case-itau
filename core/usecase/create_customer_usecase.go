package usecase

import (
	"context"
	"go.uber.org/zap"
	"transfer-api/core/domain"
	errors2 "transfer-api/core/errors"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type CreateCustomerUseCase struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
}

func NewCreateCustomerUseCase(customerRepo repository.CustomerRepository, accountRepo repository.AccountRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		CustomerRepo: customerRepo,
		AccountRepo:  accountRepo,
	}
}

func (c *CreateCustomerUseCase) Execute(ctx context.Context, input input.CreateCustomerInput) (*output.CreateCustomerOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start create customer use case")
	defer log.Info("end create customer use case")

	if err := input.Validate(); err != nil {
		err := errors2.NewValidationError(err...)
		log.Error("error while validating dto", zap.Error(err))
		return nil, err
	}

	customer := domain.NewCustomer(input.Id, input.Name, input.AccountId)

	savedCustomer, err := c.CustomerRepo.Save(*customer)

	if err != nil {
		log.Error("error while saving customer", zap.Error(err))
		return nil, err
	}

	log.Debug("customer entity saved", zap.Any("customer", savedCustomer))

	balance := util.FloatToCents(input.Balance)

	account := domain.NewAccount(input.AccountId, savedCustomer.Id, balance)

	savedAccount, err := c.AccountRepo.Save(*account)

	if err != nil {
		log.Error("error while saving account", zap.Error(err))
		return nil, err
	}

	log.Debug("account entity saved", zap.Any("account", savedAccount))

	newUser := output.NewCreateCustomerOutput(savedAccount.CustomerId, savedAccount.Id)

	return newUser, nil
}
