package repository

import (
	"errors"
	"sync"
	"transfer-api/core/domain"
)

var (
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrCustomerAlreadyExists = errors.New("customer already exists")
)

type InMemCustomerRepository struct {
	mu        sync.RWMutex
	Customers map[string]*domain.Customer
}

func NewInMemCustomerRepository() *InMemCustomerRepository {
	return &InMemCustomerRepository{
		Customers: make(map[string]*domain.Customer),
	}
}

func (r *InMemCustomerRepository) Save(customer domain.Customer) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Customers[customer.Id]; exists {
		return nil, ErrCustomerAlreadyExists
	}

	r.Customers[customer.Id] = &customer
	return &customer, nil
}

func (r *InMemCustomerRepository) GetById(id string) (*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.Customers[id]
	if !exists {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}

func (r *InMemCustomerRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Customers[id]; !exists {
		return ErrCustomerNotFound
	}

	delete(r.Customers, id)
	return nil
}

func (r *InMemCustomerRepository) GetAll() ([]*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customers := make([]*domain.Customer, 0, len(r.Customers))
	for _, customer := range r.Customers {
		customers = append(customers, customer)
	}

	return customers, nil
}
