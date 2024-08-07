package repository

import "transfer-api/src/core/domain"

type TransactionRepository interface {
	GetTransactions(accountId string) ([]domain.Transaction, error)
	SaveTransaction(transaction domain.Transaction) (domain.Transaction, error)
}
