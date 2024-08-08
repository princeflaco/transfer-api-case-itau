package domain

import (
	"transfer-api/core/errors"
)

type Account struct {
	Id         string
	CustomerId string
	Balance    int
}

func NewAccount(id string, customerId string, balance int) *Account {
	return &Account{
		Id:         id,
		Balance:    balance,
		CustomerId: customerId,
	}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		missingAmount := amount - a.Balance
		return errors.NewInsufficientFundsError(a.Id, missingAmount)
	}
	a.Balance -= amount
	return nil
}
