package domain_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transfer-api/core/domain"
	errorsCore "transfer-api/core/errors"
)

func TestNewAccount(t *testing.T) {
	expectedAccount := domain.Account{
		Id:         uuid.NewString(),
		CustomerId: uuid.NewString(),
		Balance:    100,
	}
	account := domain.NewAccount(expectedAccount.Id, expectedAccount.CustomerId, expectedAccount.Balance)
	assert.Equal(t, expectedAccount.Id, account.Id)
	assert.Equal(t, expectedAccount.CustomerId, account.CustomerId)
	assert.Equal(t, expectedAccount.Balance, account.Balance)
}

func TestAccount_Deposit(t *testing.T) {
	account := domain.NewAccount("account123", "customer456", 1000)

	account.Deposit(500)
	assert.Equal(t, 1500, account.Balance)
}

func TestAccount_Withdraw_Success(t *testing.T) {
	account := domain.NewAccount("account123", "customer456", 1000)

	err := account.Withdraw(500)
	if assert.NoError(t, err) {
		assert.Equal(t, 500, account.Balance)
	}
}

func TestAccount_Withdraw_InsufficientFunds(t *testing.T) {
	account := domain.NewAccount("account123", "customer456", 1000)

	err := account.Withdraw(1500)
	if assert.Error(t, err) {
		assert.Equal(t, 1000, account.Balance)
	}

	if assert.Error(t, err) {
		var insufficientFundsError *errorsCore.InsufficientFundsError
		ok := errors.As(err, &insufficientFundsError)
		if assert.True(t, ok) {
			assert.Equal(t, 500, insufficientFundsError.Amount)
		}
	}
}
