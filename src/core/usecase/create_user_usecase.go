package usecase

import (
	"context"
	"transfer-api/src/core/domain"
	"transfer-api/src/core/errors"
	"transfer-api/src/core/repository"
	"transfer-api/src/core/usecase/input"
	"transfer-api/src/core/usecase/output"
)

type CreateUserUseCase struct {
	UserRepo    repository.UserRepository
	AccountRepo repository.AccountRepository
}

func (c *CreateUserUseCase) Execute(ctx context.Context, input input.CreateUserInput) (*output.CreateUserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	exists, err := c.UserRepo.Exists(input.Id)

	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewDuplicatedIdError(input.Id)
	}

	user := domain.NewUser(input.Id, input.Name)

	savedUser, err := c.UserRepo.Save(*user)

	if err != nil {
		return nil, err
	}

	exists, err = c.AccountRepo.Exists(input.AccountNumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewDuplicatedIdError(input.AccountNumber)
	}

	account := domain.NewAccount(input.AccountNumber, savedUser.Id, input.Balance)

	savedAccount, err := c.AccountRepo.Save(*account)

	if err != nil {
		return nil, err
	}

	newUser := output.NewCreateUserOutput(savedAccount.UserId, savedAccount.Id)

	return newUser, nil
}
