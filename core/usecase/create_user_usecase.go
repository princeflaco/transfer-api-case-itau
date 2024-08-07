package usecase

import (
	"transfer-api/core/domain"
	"transfer-api/core/repository"
	"transfer-api/core/usecase/input"
	"transfer-api/core/usecase/output"
)

type CreateUserUseCase struct {
	UserRepo    repository.UserRepository
	AccountRepo repository.AccountRepository
}

func (c *CreateUserUseCase) Execute(input input.CreateUserInput) (*output.CreateUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	user := domain.NewUser(input.Id, input.Name)

	savedUser, err := c.UserRepo.Save(*user)

	if err != nil {
		return nil, err
	}

	account := domain.NewAccount(input.AccountId, savedUser.Id, input.Balance)

	savedAccount, err := c.AccountRepo.Save(account)

	if err != nil {
		return nil, err
	}

	newUser := output.NewCreateUserOutput(savedAccount.UserId, savedAccount.Id)

	return newUser, nil
}
