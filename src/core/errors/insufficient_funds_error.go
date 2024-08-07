package errors

import (
	"fmt"
	"transfer-api/src/core/domain"
)

type InsufficientFundsError struct {
	Account domain.Account `json:"account"`
	Amount  int            `json:"amount"`
}

var _ error = (*InsufficientFundsError)(nil)

func NewInsufficientFundsError(account domain.Account, amount int) *InsufficientFundsError {
	return &InsufficientFundsError{
		Account: account,
		Amount:  amount,
	}
}

func (err *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds from: %s, amount: %d, ", err.Account, err.Amount)
}
