package errors

import (
	"fmt"
)

type InsufficientFundsError struct {
	AccountId string `json:"accountId"`
	Amount    int    `json:"amount"`
}

var _ error = (*InsufficientFundsError)(nil)

func NewInsufficientFundsError(accountId string, amount int) *InsufficientFundsError {
	return &InsufficientFundsError{
		AccountId: accountId,
		Amount:    amount,
	}
}

func (err *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds from: %s, missing amount: %d, ", err.AccountId, err.Amount)
}
