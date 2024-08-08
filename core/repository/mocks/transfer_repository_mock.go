package mocks

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type TransferRepositoryMock struct {
	mock.Mock
}

func (t *TransferRepositoryMock) GetAll(accountId string) ([]*domain.Transfer, error) {
	r := t.Called(accountId)
	if r.Get(0) == nil {
		return nil, r.Error(1)
	}
	return r.Get(0).([]*domain.Transfer), r.Error(1)
}

func (t *TransferRepositoryMock) Save(transfer domain.Transfer) (*domain.Transfer, error) {
	r := t.Called(transfer)
	if r.Get(0) == nil {
		return nil, r.Error(1)
	}
	return r.Get(0).(*domain.Transfer), r.Error(1)
}
