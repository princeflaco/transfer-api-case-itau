package repository

import "transfer-api/core/domain"

type CustomerRepository interface {
	GetById(id string) (*domain.Customer, error)
	GetAll() ([]*domain.Customer, error)
	Save(customer domain.Customer) (*domain.Customer, error)
	Delete(id string) error
}
