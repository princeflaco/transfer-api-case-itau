package usecase

import (
	"transfer-api/core/repository"
	"transfer-api/core/usecase/output"
)

type GetCustomerByAccountIdUseCase struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
}

func NewGetCustomerByAccountIdUseCase(customerRepo repository.CustomerRepository, accountRepo repository.AccountRepository) *GetCustomerByAccountIdUseCase {
	return &GetCustomerByAccountIdUseCase{
		CustomerRepo: customerRepo,
		AccountRepo:  accountRepo,
	}
}

func (uc *GetCustomerByAccountIdUseCase) Execute(accountId string) (*output.GetCustomerByAccountIdOutput, error) {
	account, err := uc.AccountRepo.GetById(accountId)
	if err != nil {
		return nil, err
	}
	customer, err := uc.CustomerRepo.GetById(account.CustomerId)
	if err != nil {
		return nil, err
	}
	balance := centsToFloat64(account.Balance)
	customerOutput := output.NewGetCustomerByAccountIdOutput(customer.Id, customer.Name, account.Id, balance)
	return customerOutput, nil
}

func centsToFloat64(cents int) float64 {
	return float64(cents) / 100.0
}
