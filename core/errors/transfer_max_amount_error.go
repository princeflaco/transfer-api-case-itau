package errors

import (
	"fmt"
	"strconv"
)

type TransferMaxAmountError struct {
	Amount float64
}

func NewTransferMaxAmountError(amount float64) *TransferMaxAmountError {
	return &TransferMaxAmountError{
		Amount: amount,
	}
}

func (e *TransferMaxAmountError) Error() string {
	amount := strconv.FormatFloat(e.Amount, 'f', 1, 64)
	return fmt.Sprintf("Amount of %s exceeds the maximum amount set", amount)
}
