package repository

import "transfer-api/src/core/domain"

type AccountRepository interface {
	Save(domain.Account) (*domain.Account, error)
	GetById(id string) (*domain.Account, error)
	Exists(id string) (bool, error)
	Delete(id string) error
}
