package integration

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"transfer-api/adapter/repository"
	"transfer-api/core/domain"
	"transfer-api/core/usecase"
	"transfer-api/core/usecase/input"
	"transfer-api/core/util"
)

const (
	TransferCount  = 100
	TransferAmount = 10.0
)

func setupInMemRepositories() (*repository.InMemAccountRepository, *repository.InMemTransferRepository) {
	accountRepo := repository.NewInMemAccountRepository()
	transferRepo := repository.NewInMemTransferRepository()
	return accountRepo, transferRepo
}

func TestCreateTransferUseCase_Concurrent_Execute(t *testing.T) {
	accountRepo, transferRepo := setupInMemRepositories()
	useCase := usecase.NewCreateTransferUseCase(transferRepo, accountRepo)

	accountFromId := uuid.NewString()
	accountToId := uuid.NewString()

	accountFrom := &domain.Account{
		Id:         accountFromId,
		CustomerId: uuid.NewString(),
		Balance:    100000,
	}
	accountTo := &domain.Account{
		Id:         accountToId,
		CustomerId: uuid.NewString(),
		Balance:    50000,
	}

	_, err := accountRepo.Save(*accountFrom)
	if err != nil {
		t.Error(err)
	}

	_, err = accountRepo.Save(*accountTo)
	if err != nil {
		t.Error(err)
	}

	var wg sync.WaitGroup
	wg.Add(TransferCount)

	for i := 0; i < TransferCount; i++ {
		go func() {
			defer wg.Done()

			ctx := util.NewTestContext()
			transferInput := input.TransferInput{
				TargetAccountId: accountToId,
				Amount:          TransferAmount,
			}

			_, err = useCase.Execute(ctx, transferInput, accountFromId)
			if err != nil {
				t.Error(err)
			}
		}()
	}

	wg.Wait()

	accountFrom, err = accountRepo.GetById(accountFromId)
	if err != nil {
		t.Error(err)
	}

	accountTo, err = accountRepo.GetById(accountToId)
	if err != nil {
		t.Error(err)
	}

	amount := TransferCount * util.FloatToCents(TransferAmount)

	expectedFromBalance := 100000 - amount
	expectedToBalance := 50000 + amount

	assert.Equal(t, expectedFromBalance, accountFrom.GetBalance(), "Unexpected balance for accountFrom")
	assert.Equal(t, expectedToBalance, accountTo.GetBalance(), "Unexpected balance for accountTo")
}
