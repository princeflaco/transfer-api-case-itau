package usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"transfer-api/core/domain"
	"transfer-api/core/repository/mocks"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/output"
	"transfer-api/core/util"
)

func TestGetTransferHistoryUseCase_Success(t *testing.T) {
	mockTransferRepo := new(mocks.TransferRepositoryMock)

	accountId := uuid.NewString()
	now := time.Now()
	nowString := now.Format(domain.TimeFormat)
	nowPlusHour := now.Add(time.Hour).Format(domain.TimeFormat)
	nowPlus2Hours := now.Add(2 * time.Hour).Format(domain.TimeFormat)

	mockTransfers := []*domain.Transfer{
		{Id: uuid.NewString(), AccountId: accountId, TargetAccountId: uuid.NewString(), Amount: 100100, Date: nowString, Success: true},
		{Id: uuid.NewString(), AccountId: accountId, TargetAccountId: uuid.NewString(), Amount: 100200, Date: nowPlusHour, Success: false},
		{Id: uuid.NewString(), AccountId: accountId, TargetAccountId: uuid.NewString(), Amount: 100300, Date: nowPlus2Hours, Success: true},
	}

	assert.Len(t, mockTransfers, 3)

	transferHistory := []output.TransferHistoryOutput{
		{Id: mockTransfers[2].Id, TargetAccountId: mockTransfers[2].TargetAccountId, Amount: util.CentsToFloat64(mockTransfers[2].Amount), Date: mockTransfers[2].Date, Success: mockTransfers[2].Success},
		{Id: mockTransfers[1].Id, TargetAccountId: mockTransfers[1].TargetAccountId, Amount: util.CentsToFloat64(mockTransfers[1].Amount), Date: mockTransfers[1].Date, Success: mockTransfers[1].Success},
		{Id: mockTransfers[0].Id, TargetAccountId: mockTransfers[0].TargetAccountId, Amount: util.CentsToFloat64(mockTransfers[0].Amount), Date: mockTransfers[0].Date, Success: mockTransfers[0].Success},
	}

	assert.Len(t, transferHistory, 3)

	mockTransferRepo.On("GetAll", accountId).Return(mockTransfers, nil)

	uc := usecase.NewGetTransferHistoryUseCase(mockTransferRepo)

	ucOutput, err := uc.Execute(accountId)

	assert.NoError(t, err)
	assert.Len(t, transferHistory, len(mockTransfers))
	assert.Equal(t, transferHistory, ucOutput)

	mockTransferRepo.AssertExpectations(t)
}
