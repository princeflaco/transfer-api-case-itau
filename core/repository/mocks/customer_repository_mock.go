package mocks

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (u *CustomerRepositoryMock) GetById(id string) (domain.Customer, error) {
	r := u.Called(id)
	return r.Get(0).(domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) GetAll() ([]domain.Customer, error) {
	r := u.Called()
	return r.Get(0).([]domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) Save(customer domain.Customer) (domain.Customer, error) {
	r := u.Called(customer)
	return r.Get(0).(domain.Customer), r.Error(1)
}

func (u *CustomerRepositoryMock) Delete(id string) error {
	r := u.Called(id)
	return r.Error(0)
}
