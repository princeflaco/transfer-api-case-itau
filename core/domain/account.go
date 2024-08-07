package domain

import (
	"github.com/google/uuid"
	"sync"
	"transfer-api/core/errors"
)

type Account struct {
	Id         string
	CustomerId string
	Balance    int
	mu         sync.Mutex
}

func NewAccount(id string, customerId string, balance int) *Account {
	if id == "" {
		id = uuid.NewString()
	}
	return &Account{
		Id:         id,
		Balance:    balance,
		CustomerId: customerId,
	}
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount < a.Balance {
		missingAmount := amount - a.Balance
		return errors.NewInsufficientFundsError(a.Id, missingAmount)
	}
	a.Balance -= amount
	return nil
}
