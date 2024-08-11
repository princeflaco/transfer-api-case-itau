package repository

import "transfer-api/core/domain"

type AccountRepository interface {
	Save(domain.Account) (*domain.Account, error)
	GetById(id string) (*domain.Account, error)
	Update(domain.Account) (*domain.Account, error)
	Delete(id string) error
}
