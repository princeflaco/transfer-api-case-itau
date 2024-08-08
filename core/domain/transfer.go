package domain

import (
	"github.com/google/uuid"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const TransferMaxAmount = 10000.0

type Transfer struct {
	Id              string
	AccountId       string
	TargetAccountId string
	Success         bool
	Amount          int
	Date            string
}

func NewTransfer(accountId string, targetAccountId string, amount int) *Transfer {
	return &Transfer{
		AccountId:       accountId,
		TargetAccountId: targetAccountId,
		Amount:          amount,
		Date:            time.Now().Format(TimeFormat),
		Id:              uuid.NewString(),
	}
}

func (t *Transfer) Successful(successful bool) {
	t.Success = successful
}
