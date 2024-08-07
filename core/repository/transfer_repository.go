package repository

import "transfer-api/core/domain"

type TransferRepository interface {
	GetTransfers(accountId string) ([]*domain.Transfer, error)
	SaveTransfer(transfer domain.Transfer) (*domain.Transfer, error)
}
