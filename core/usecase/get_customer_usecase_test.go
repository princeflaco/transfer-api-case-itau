package usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"transfer-api/core/domain"
	"transfer-api/core/errors"
	"transfer-api/core/usecase"
	"transfer-api/core/util"
)

func TestGetCustomerUseCase_Success(t *testing.T) {
	mockCustomerRepo, mockAccountRepo, _, ctx := setupMockDependencies()
	uc := usecase.NewGetCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	customer := domain.Customer{
		Id:        uuid.NewString(),
		Name:      "Leonardo",
		AccountId: uuid.NewString(),
	}

	account := domain.Account{
		Id:         customer.AccountId,
		CustomerId: customer.Id,
		Balance:    1000,
	}

	mockAccountRepo.On("GetById", mock.Anything).Return(&account, nil)
	mockCustomerRepo.On("GetById", mock.Anything).Return(&customer, nil)

	customerOutput, err := uc.Execute(ctx, account.Id)

	assert.NoError(t, err)
	assert.Equal(t, customer.Id, customerOutput.Id)
	assert.Equal(t, customer.Name, customerOutput.Name)
	assert.Equal(t, account.Id, customerOutput.AccountNumber)
	assert.Equal(t, util.CentsToFloat64(account.Balance), customerOutput.Balance)

	mockAccountRepo.AssertExpectations(t)
	mockCustomerRepo.AssertExpectations(t)
}

func TestGetCustomerUseCase_AccountRepoError(t *testing.T) {
	mockCustomerRepo, mockAccountRepo, _, ctx := setupMockDependencies()
	uc := usecase.NewGetCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	accountId := uuid.NewString()

	notFoundErr := errors.NewNotFoundError("account", accountId)

	mockAccountRepo.On("GetById", mock.Anything).Return(nil, notFoundErr)

	customerOutput, err := uc.Execute(ctx, mock.Anything)

	if assert.Error(t, err, notFoundErr) {
		assert.Nil(t, customerOutput)
	}

	mockAccountRepo.AssertExpectations(t)
	mockCustomerRepo.AssertExpectations(t)
}

func TestGetCustomerUseCase_CustomerRepoError(t *testing.T) {
	mockCustomerRepo, mockAccountRepo, _, ctx := setupMockDependencies()
	uc := usecase.NewGetCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	account := domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1000,
	}

	notFoundErr := errors.NewNotFoundError("customer", account.CustomerId)

	mockAccountRepo.On("GetById", mock.Anything).Return(&account, nil)
	mockCustomerRepo.On("GetById", mock.Anything).Return(nil, notFoundErr)

	customerOutput, err := uc.Execute(ctx, account.Id)

	if assert.Error(t, err, notFoundErr) {
		assert.Nil(t, customerOutput)
	}

	mockAccountRepo.AssertExpectations(t)
	mockCustomerRepo.AssertExpectations(t)
}
