package repository

import (
	"sync"
	"transfer-api/core/domain"
)

type InMemTransferRepository struct {
	mu        sync.RWMutex
	transfers map[string]*domain.Transfer
}

func NewInMemTransferRepository() *InMemTransferRepository {
	return &InMemTransferRepository{
		transfers: make(map[string]*domain.Transfer),
	}
}

func (r *InMemTransferRepository) Save(transfer domain.Transfer) (*domain.Transfer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.transfers[transfer.Id] = &transfer
	return &transfer, nil
}

func (r *InMemTransferRepository) GetAll(accountId string) ([]*domain.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transfers := make([]*domain.Transfer, 0, len(r.transfers))
	for _, transfer := range r.transfers {
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
