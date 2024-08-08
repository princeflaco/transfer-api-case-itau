package domain_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transfer-api/core/domain"
)

func TestNewCustomer(t *testing.T) {
	expectedCustomer := domain.Customer{
		Id:        uuid.NewString(),
		Name:      "leonardo",
		AccountId: uuid.NewString(),
	}
	customer := domain.NewCustomer(expectedCustomer.Id, expectedCustomer.Name, expectedCustomer.AccountId)
	assert.Equal(t, expectedCustomer.Id, customer.Id)
	assert.Equal(t, expectedCustomer.Name, customer.Name)
	assert.Equal(t, expectedCustomer.AccountId, customer.AccountId)
}
