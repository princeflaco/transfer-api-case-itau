package repository

import "transfer-api/core/domain"

type TransferRepository interface {
	GetAll(accountId string) ([]*domain.Transfer, error)
	Save(transfer domain.Transfer) (*domain.Transfer, error)
}
