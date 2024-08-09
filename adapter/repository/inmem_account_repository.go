package repository

import (
	"errors"
	"sync"
	"transfer-api/core/domain"
)

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")
)

type InMemAccountRepository struct {
	mu       sync.RWMutex
	Accounts map[string]*domain.Account
}

func NewInMemAccountRepository() *InMemAccountRepository {
	return &InMemAccountRepository{
		Accounts: make(map[string]*domain.Account),
	}
}

func (r *InMemAccountRepository) Save(account domain.Account) (*domain.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Accounts[account.Id]; exists {
		return nil, ErrAccountAlreadyExists
	}

	r.Accounts[account.Id] = &account
	return &account, nil
}

func (r *InMemAccountRepository) Update(account domain.Account) (*domain.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Accounts[account.Id]; !exists {
		return nil, ErrAccountNotFound
	}
	
	r.Accounts[account.Id] = &account
	return &account, nil
}

func (r *InMemAccountRepository) GetById(id string) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, exists := r.Accounts[id]
	if !exists {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (r *InMemAccountRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Accounts[id]; !exists {
		return ErrAccountNotFound
	}

	delete(r.Accounts, id)
	return nil
}
