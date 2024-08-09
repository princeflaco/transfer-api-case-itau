package usecase

import (
	"context"
	"go.uber.org/zap"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

type GetAllCustomersUseCase struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
}

func NewGetAllCustomersUseCase(customerRepo repository.CustomerRepository, accountRepository repository.AccountRepository) *GetAllCustomersUseCase {
	return &GetAllCustomersUseCase{
		CustomerRepo: customerRepo,
		AccountRepo:  accountRepository,
	}
}

func (uc *GetAllCustomersUseCase) Execute(ctx context.Context) ([]*output.GetCustomerOutput, error) {
	requestId := util.GetRequestIDFromContext(ctx)
	log := util.GetLoggerFromContext(ctx).With(zap.String("request_id", requestId))

	log.Info("start get all customers use case")
	defer log.Info("end get all customers use case")

	customers, err := uc.CustomerRepo.GetAll()
	if err != nil {
		log.Error("error while retrieving customers from database")
		log.Error(err.Error())
		return []*output.GetCustomerOutput{}, err
	}
	log.Debug("retrieved all customers from database")

	var outputs []*output.GetCustomerOutput
	for _, customer := range customers {
		account, err := uc.AccountRepo.GetById(customer.AccountId)
		if err != nil {
			log.Debug("could not retrieve account of customer: " + customer.Id + " from database")
			continue
		}
		log.Debug("retrieved account of customer: " + customer.AccountId + " from database")
		balance := util.CentsToFloat64(account.Balance)
		customerOutput := output.NewGetCustomerOutput(customer.Id, customer.Name, account.Id, balance)
		outputs = append(outputs, customerOutput)
	}
	return outputs, nil
}
