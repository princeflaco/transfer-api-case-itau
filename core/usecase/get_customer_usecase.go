package usecase

import (
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

func (uc *GetCustomerUseCase) Execute(accountId string) (*output.GetCustomerOutput, error) {
	account, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		return nil, err
	}
	customer, err := uc.CustomerRepo.GetById(account.CustomerId)
	if err != nil {
		return nil, err
	}
	balance := util.CentsToFloat64(account.Balance)
	customerOutput := output.NewGetCustomerOutput(customer.Id, customer.Name, account.Id, balance)
	return customerOutput, nil
}
