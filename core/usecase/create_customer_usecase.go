package usecase

import (
	"transfer-api/core/domain"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
)

type CreateCustomerUseCase struct {
	UserRepo    repository.CustomerRepository
	AccountRepo repository.AccountRepository
}

func (c *CreateCustomerUseCase) Execute(input input.CreateCustomerInput) (*output.CreateCustomerOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user := domain.NewCustomer(input.Id, input.Name)

	savedUser, err := c.UserRepo.Save(*user)

	if err != nil {
		return nil, err
	}

	account := domain.NewAccount(input.AccountId, savedUser.Id, input.Balance)

	savedAccount, err := c.AccountRepo.Save(account)

	if err != nil {
		return nil, err
	}

	newUser := output.NewCreateCustomerOutput(savedAccount.CustomerId, savedAccount.Id)

	return newUser, nil
}
