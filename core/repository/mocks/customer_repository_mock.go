package mocks

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (u *CustomerRepositoryMock) GetById(id string) (*domain.Customer, error) {
	r := u.Called(id)
	if r.Get(0) == nil {
		return nil, r.Error(1)
	}
	return r.Get(0).(*domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) GetAll() ([]*domain.Customer, error) {
	r := u.Called()
	if r.Get(0) == nil {
		return nil, r.Error(1)
	}
	return r.Get(0).([]*domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) Save(customer domain.Customer) (*domain.Customer, error) {
	r := u.Called(customer)
	if r.Get(0) == nil {
		return nil, r.Error(1)
	}
	return r.Get(0).(*domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) Delete(id string) error {
	r := u.Called(id)
	return r.Error(0)
}
