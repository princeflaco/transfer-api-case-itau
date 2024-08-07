package mocks

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (a *AccountRepositoryMock) Save(account *domain.Account) (*domain.Account, error) {
	r := a.Called(&account)
	return r.Get(0).(*domain.Account), r.Error(1)
}

func (a *AccountRepositoryMock) GetById(id string) (*domain.Account, error) {
	r := a.Called(id)
	return r.Get(0).(*domain.Account), r.Error(1)
}

func (a *AccountRepositoryMock) Delete(id string) error {
	r := a.Called(id)
	return r.Error(0)
}
