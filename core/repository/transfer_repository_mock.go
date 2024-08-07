package repository

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type TransferRepositoryMock struct {
	mock.Mock
}

func (t *TransferRepositoryMock) GetTransfers(accountId string) ([]domain.Transfer, error) {
	r := t.Called(accountId)
	return r.Get(0).([]domain.Transfer), r.Error(1)
}

func (t *TransferRepositoryMock) SaveTransfer(transfer domain.Transfer) (domain.Transfer, error) {
	r := t.Called(transfer)
	return r.Get(0).(domain.Transfer), r.Error(1)
}
