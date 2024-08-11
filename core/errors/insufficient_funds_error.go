package errors

import (
	"fmt"
	"transfer-api/core/util"
)

type InsufficientFundsError struct {
	Amount int `json:"amount"`
}

var _ error = (*InsufficientFundsError)(nil)

func NewInsufficientFundsError(amount int) *InsufficientFundsError {
	return &InsufficientFundsError{
		Amount: amount,
	}
}

func (err *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds, missing amount: %f", util.CentsToFloat64(err.Amount))
}
