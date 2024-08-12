package usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"transfer-api/core/domain"
	errors2 "transfer-api/core/errors"
	"transfer-api/core/service"
	"transfer-api/core/service/input"
	"transfer-api/core/usecase"
	"transfer-api/core/util"
)

func TestCreateTransferUseCase_Execute_Success(t *testing.T) {
	_, mockAccountRepo, mockTransferRepo, ctx := setupMockDependencies()
	svc := service.NewTransferServiceImpl(mockTransferRepo, mockAccountRepo)
	useCase := usecase.NewCreateTransferUseCase(svc)

	accountFrom := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1500,
	}

	accountTo := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1000,
	}

	i := input.TransferInput{
		AccountId:       accountFrom.Id,
		TargetAccountId: accountTo.Id,
		Amount:          5.0,
	}

	expectedTransfer := domain.Transfer{
		Id:              uuid.NewString(),
		AccountId:       accountFrom.Id,
		TargetAccountId: accountTo.Id,
		Success:         true,
		Amount:          util.FloatToCents(i.Amount),
		Date:            time.Now().Format(domain.TimeFormat),
	}

	mockAccountRepo.On("GetById", accountFrom.Id).Return(accountFrom, nil)
	mockAccountRepo.On("GetById", accountTo.Id).Return(accountTo, nil)

	mockTransferRepo.On("Save", mock.Anything).Return(&expectedTransfer, nil)
	mockAccountRepo.On("Update", mock.Anything).Return(accountFrom, nil)
	mockAccountRepo.On("Update", mock.Anything).Return(accountTo, nil)

	output, err := useCase.Execute(ctx, i, accountFrom.Id)

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, expectedTransfer.Id, output.Id)
	assert.True(t, output.Successful)
	assert.Equal(t, expectedTransfer.Date, output.Timestamp)

	mockAccountRepo.AssertExpectations(t)
	mockTransferRepo.AssertExpectations(t)
}

func TestCreateTransferUseCase_Execute_InsufficientFunds(t *testing.T) {
	_, mockAccountRepo, mockTransferRepo, ctx := setupMockDependencies()
	svc := service.NewTransferServiceImpl(mockTransferRepo, mockAccountRepo)
	useCase := usecase.NewCreateTransferUseCase(svc)

	accountFrom := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    500,
	}
	accountTo := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1000,
	}

	i := input.TransferInput{
		AccountId:       accountFrom.Id,
		TargetAccountId: accountTo.Id,
		Amount:          600.0,
	}

	expectedTransfer := domain.Transfer{
		Id:              uuid.NewString(),
		AccountId:       accountFrom.Id,
		TargetAccountId: accountTo.Id,
		Success:         false,
		Amount:          util.FloatToCents(i.Amount),
		Date:            time.Now().Format(domain.TimeFormat),
	}

	mockAccountRepo.On("GetById", accountFrom.Id).Return(accountFrom, nil)
	mockAccountRepo.On("GetById", accountTo.Id).Return(accountTo, nil)
	mockTransferRepo.On("Save", mock.Anything).Return(&expectedTransfer, nil)

	output, err := useCase.Execute(ctx, i, accountFrom.Id)

	assert.NoError(t, err)

	if assert.NotNil(t, output) {
		assert.Equal(t, expectedTransfer.Id, output.Id)
		assert.Equal(t, expectedTransfer.Success, output.Successful)
		assert.Equal(t, expectedTransfer.Date, output.Timestamp)
	}

	mockAccountRepo.AssertExpectations(t)
	mockTransferRepo.AssertExpectations(t)
}

func TestCreateTransferUseCase_Execute_RepoError(t *testing.T) {
	_, mockAccountRepo, mockTransferRepo, ctx := setupMockDependencies()
	svc := service.NewTransferServiceImpl(mockTransferRepo, mockAccountRepo)
	useCase := usecase.NewCreateTransferUseCase(svc)

	accountFrom := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1500,
	}

	accountTo := &domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    1000,
	}

	i := input.TransferInput{
		AccountId:       accountFrom.Id,
		TargetAccountId: accountTo.Id,
		Amount:          5.0,
	}

	notFoundErr := errors2.NewNotFoundError("account", accountFrom.Id)

	mockAccountRepo.On("GetById", accountFrom.Id).Return(nil, notFoundErr)

	output, err := useCase.Execute(ctx, i, accountFrom.Id)

	assert.Error(t, err)
	assert.Nil(t, output)
	mockAccountRepo.AssertExpectations(t)
}
