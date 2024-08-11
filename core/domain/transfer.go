package domain

import (
	"github.com/google/uuid"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type Transfer struct {
	Id              string
	AccountId       string
	TargetAccountId string
	Success         bool
	Reason          string
	Amount          int
	Date            string
}

func NewTransfer(accountId string, targetAccountId string, amount int) *Transfer {
	return &Transfer{
		AccountId:       accountId,
		TargetAccountId: targetAccountId,
		Amount:          amount,
		Success:         true,
		Date:            time.Now().Format(TimeFormat),
		Id:              uuid.NewString(),
	}
}

func (t *Transfer) NotSuccessful(reason string) {
	t.Success = false
	t.Reason = reason
}
