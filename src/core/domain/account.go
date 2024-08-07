package domain

import (
	"github.com/google/uuid"
	"transfer-api/src/core/errors"
)

type Account struct {
	Id      string
	UserId  string
	Balance int
}

func NewAccount(id string, userId string, balance int) *Account {
	if id == "" {
		id = uuid.NewString()
	}
	return &Account{
		Id:      id,
		Balance: balance,
		UserId:  userId,
	}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount < a.Balance {
		missingAmount := amount - a.Balance
		return errors.NewInsufficientFundsError(*a, missingAmount)
	}
	a.Balance -= amount
	return nil
}
