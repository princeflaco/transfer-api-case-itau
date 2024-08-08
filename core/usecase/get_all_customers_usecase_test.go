package usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transfer-api/core/domain"
	"transfer-api/core/repository/mocks"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

func TestGetAllCustomersUseCase_Execute_Success(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepositoryMock)
	mockAccountRepo := new(mocks.AccountRepositoryMock)

	useCase := usecase.NewGetAllCustomersUseCase(mockCustomerRepo, mockAccountRepo)

	leonardo := domain.NewCustomer(uuid.NewString(), "Leonardo", uuid.NewString())
	lucas := domain.NewCustomer(uuid.NewString(), "Lucas", uuid.NewString())
	joao := domain.NewCustomer(uuid.NewString(), "Joao", uuid.NewString())

	customers := []*domain.Customer{leonardo, lucas, joao}

	account1 := domain.NewAccount(leonardo.AccountId, leonardo.Id, 1000)
	account2 := domain.NewAccount(lucas.AccountId, lucas.Id, 1000)
	account3 := domain.NewAccount(joao.AccountId, joao.Id, 1000)

	accounts := []*domain.Account{account1, account2, account3}

	balance1 := util.CentsToFloat64(account1.Balance)
	output1 := output.NewGetCustomerOutput(leonardo.Id, leonardo.Name, leonardo.AccountId, balance1)

	balance2 := util.CentsToFloat64(account2.Balance)
	output2 := output.NewGetCustomerOutput(lucas.Id, lucas.Name, lucas.AccountId, balance2)

	balance3 := util.CentsToFloat64(account3.Balance)
	output3 := output.NewGetCustomerOutput(joao.Id, joao.Name, joao.AccountId, balance3)

	expectedOutputs := []*output.GetCustomerOutput{output1, output2, output3}

	mockCustomerRepo.On("GetAll").Return(customers, nil)

	for _, account := range accounts {
		mockAccountRepo.On("GetById", account.Id).Return(account, nil)
	}

	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Equal(t, expectedOutputs, result)

	mockCustomerRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}
