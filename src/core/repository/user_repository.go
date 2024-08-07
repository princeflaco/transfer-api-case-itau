package repository

import "transfer-api/src/core/domain"

type UserRepository interface {
	GetById(userId string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Save(user domain.User) (domain.User, error)
	Delete(userId string) error
}
