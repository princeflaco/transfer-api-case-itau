package domain_test

import (
	"testing"
	"time"
	"transfer-api/core/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	accountId := uuid.NewString()
	targetAccountId := uuid.NewString()
	amount := 500

	transfer := domain.NewTransfer(accountId, targetAccountId, amount)

	assert.Equal(t, accountId, transfer.AccountId)
	assert.Equal(t, targetAccountId, transfer.TargetAccountId)
	assert.Equal(t, amount, transfer.Amount)
	assert.True(t, transfer.Success)

	_, err := uuid.Parse(transfer.Id)
	assert.NoError(t, err)

	_, err = time.Parse(domain.TimeFormat, transfer.Date)
	assert.NoError(t, err)
}

func TestNewTransfer_NotSuccessful(t *testing.T) {
	accountId := uuid.NewString()
	targetAccountId := uuid.NewString()
	amount := 500

	transfer := domain.NewTransfer(accountId, targetAccountId, amount)
	transfer.NotSuccessful("test")

	assert.Equal(t, accountId, transfer.AccountId)
	assert.Equal(t, targetAccountId, transfer.TargetAccountId)
	assert.Equal(t, amount, transfer.Amount)
	assert.False(t, transfer.Success)
	assert.Equal(t, "test", transfer.Reason)

	_, err := uuid.Parse(transfer.Id)
	assert.NoError(t, err)

	_, err = time.Parse(domain.TimeFormat, transfer.Date)
	assert.NoError(t, err)
}
