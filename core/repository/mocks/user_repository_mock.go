package mocks

import (
	"github.com/stretchr/testify/mock"
	"transfer-api/core/domain"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) GetById(id string) (domain.User, error) {
	r := u.Called(id)
	return r.Get(0).(domain.User), r.Error(1)
}

func (u *UserRepositoryMock) GetAll() ([]domain.User, error) {
	r := u.Called()
	return r.Get(0).([]domain.User), r.Error(1)
}

func (u *UserRepositoryMock) Save(user domain.User) (domain.User, error) {
	r := u.Called(user)
	return r.Get(0).(domain.User), r.Error(1)
}

func (u *UserRepositoryMock) Delete(id string) error {
	r := u.Called(id)
	return r.Error(0)
}
