package usecase

import (
	"transfer-api/core/repository"
	"transfer-api/core/usecase/output"
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
	balance := CentsToFloat64(account.Balance)
	customerOutput := output.NewGetCustomerOutput(customer.Id, customer.Name, account.Id, balance)
	return customerOutput, nil
}

func CentsToFloat64(cents int) float64 {
	return float64(cents) / 100.0
}
