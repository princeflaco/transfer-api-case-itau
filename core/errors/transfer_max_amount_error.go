package errors

import (
	"fmt"
	"strconv"
	"transfer-api/core/domain"
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
	maximumAmount := strconv.FormatFloat(domain.TransferMaxAmount, 'f', 10, 64)
	amount := strconv.FormatFloat(e.Amount, 'f', 10, 64)
	return fmt.Sprintf("Amount of %s exceeds the maximum set of %s", amount, maximumAmount)
}
