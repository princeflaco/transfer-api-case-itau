package usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"transfer-api/core/domain"
	"transfer-api/core/errors"
	"transfer-api/core/repository/mocks"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/input"
	"transfer-api/core/util"
)

func TestCreateCustomerUseCase_Success(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	useCase := usecase.NewCreateCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	i := input.CreateCustomerInput{
		Id:        uuid.NewString(),
		Name:      "Leonardo",
		AccountId: uuid.NewString(),
		Balance:   1000.0,
	}

	expectedCustomer := domain.Customer{
		Id:        i.Id,
		Name:      i.Name,
		AccountId: i.AccountId,
	}

	expectedAccount := domain.Account{
		Id:         i.AccountId,
		CustomerId: i.Id,
		Balance:    util.FloatToCents(i.Balance),
	}

	mockCustomerRepo.On("Save", mock.Anything).Return(&expectedCustomer, nil)
	mockAccountRepo.On("Save", mock.Anything).Return(&expectedAccount, nil)

	output, err := useCase.Execute(i)

	if assert.NoError(t, err) {
		assert.Equal(t, i.Id, output.Id)
		assert.Equal(t, i.AccountId, output.AccountId)
	}

	mockCustomerRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}

func TestCreateCustomerUseCase_ValidationError(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	useCase := usecase.NewCreateCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	i := input.CreateCustomerInput{
		Id:        "",
		Name:      "",
		AccountId: "",
		Balance:   1000.0,
	}

	invalidFieldErr := errors.NewInvalidFieldError("name", "Should not be empty")

	output, err := useCase.Execute(i)

	assert.Error(t, err, invalidFieldErr)
	assert.Nil(t, output)

	mockCustomerRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}

func TestCreateCustomerUseCase_CustomerRepoError(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	useCase := usecase.NewCreateCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	i := input.CreateCustomerInput{
		Id:        uuid.NewString(),
		Name:      "Leonardo",
		AccountId: uuid.NewString(),
		Balance:   1000.0,
	}

	duplicatedErr := errors.NewDuplicatedEntityError(i.Id)

	mockCustomerRepo.On("Save", mock.Anything).Return(nil, duplicatedErr)

	output, err := useCase.Execute(i)

	assert.Error(t, err, duplicatedErr)
	assert.Nil(t, output)

	mockCustomerRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}

func TestCreateCustomerUseCase_AccountRepoError(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	useCase := usecase.NewCreateCustomerUseCase(mockCustomerRepo, mockAccountRepo)

	i := input.CreateCustomerInput{
		Id:        uuid.NewString(),
		Name:      "Leonardo",
		AccountId: uuid.NewString(),
		Balance:   1000.0,
	}

	expectedCustomer := domain.Customer{
		Id:        i.Id,
		Name:      i.Name,
		AccountId: i.AccountId,
	}

	duplicatedErr := errors.NewDuplicatedEntityError(i.Id)

	mockCustomerRepo.On("Save", mock.Anything).Return(&expectedCustomer, nil)
	mockAccountRepo.On("Save", mock.Anything).Return(nil, duplicatedErr)

	output, err := useCase.Execute(i)

	assert.Error(t, err, duplicatedErr)
	assert.Nil(t, output)

	mockCustomerRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}
