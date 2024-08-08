package usecase

import (
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

func (uc *GetAllCustomersUseCase) Execute() ([]*output.GetCustomerOutput, error) {
	customers, err := uc.CustomerRepo.GetAll()
	if err != nil {
		return []*output.GetCustomerOutput{}, err
	}
	var outputs []*output.GetCustomerOutput
	for _, customer := range customers {
		account, err := uc.AccountRepo.GetById(customer.AccountId)
		if err != nil {
			continue
		}
		balance := util.CentsToFloat64(account.Balance)
		customerOutput := output.NewGetCustomerOutput(customer.Id, customer.Name, account.Id, balance)
		outputs = append(outputs, customerOutput)
	}
	return outputs, nil
}
