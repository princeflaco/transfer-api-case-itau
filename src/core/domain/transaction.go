package domain

import (
	"github.com/google/uuid"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type Transaction struct {
	Id        string
	AccountId int
	Amount    int
	Date      string
}

func NewTransaction(accountId, amount int) *Transaction {
	return &Transaction{
		AccountId: accountId,
		Amount:    amount,
		Date:      time.Now().Format(TimeFormat),
		Id:        uuid.NewString(),
	}
}
