package repository

import "transfer-api/src/core/domain"

type AccountRepository interface {
	Save(domain.Account) (*domain.Account, error)
	GetById(id int) (*domain.Account, error)
	Delete(id int) error
}
