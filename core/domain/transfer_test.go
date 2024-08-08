package domain_test

import (
	"testing"
	"time"
	"transfer-api/core/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	accountId := "account123"
	targetAccountId := "targetAccount456"
	amount := 500

	transfer := domain.NewTransfer(accountId, targetAccountId, amount)

	assert.Equal(t, accountId, transfer.AccountId)
	assert.Equal(t, targetAccountId, transfer.TargetAccountId)
	assert.Equal(t, amount, transfer.Amount)
	assert.False(t, transfer.Success)

	_, err := uuid.Parse(transfer.Id)
	assert.NoError(t, err)

	_, err = time.Parse(domain.TimeFormat, transfer.Date)
	assert.NoError(t, err)
}

func TestTransfer_Successful(t *testing.T) {
	accountId := "account123"
	targetAccountId := "targetAccount456"
	amount := 500

	transfer := domain.NewTransfer(accountId, targetAccountId, amount)

	transfer.Successful(true)
	assert.True(t, transfer.Success)

	transfer.Successful(false)
	assert.False(t, transfer.Success)
}
