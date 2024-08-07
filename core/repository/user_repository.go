package repository

import "transfer-api/core/domain"

type UserRepository interface {
	GetById(id string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Save(user domain.User) (domain.User, error)
	Delete(id string) error
}
