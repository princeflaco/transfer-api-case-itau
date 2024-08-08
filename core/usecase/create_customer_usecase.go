package usecase

import (
	"transfer-api/core/domain"
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

func (c *CreateCustomerUseCase) Execute(input input.CreateCustomerInput) (*output.CreateCustomerOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user := domain.NewCustomer(input.Id, input.Name, input.AccountId)

	savedUser, err := c.CustomerRepo.Save(*user)

	if err != nil {
		return nil, err
	}

	balance := util.FloatToCents(input.Balance)

	account := domain.NewAccount(input.AccountId, savedUser.Id, balance)

	savedAccount, err := c.AccountRepo.Save(*account)

	if err != nil {
		return nil, err
	}

	newUser := output.NewCreateCustomerOutput(savedAccount.CustomerId, savedAccount.Id)

	return newUser, nil
}
