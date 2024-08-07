package repository

import "transfer-api/src/core/domain"

type UserRepository interface {
	GetById(id string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Save(user domain.User) (domain.User, error)
	Delete(id string) error
	Exists(id string) (bool, error)
}
