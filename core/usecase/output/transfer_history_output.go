package output

type TransferHistoryOutput struct {
	Id              string  `json:"id"`
	AccountId       string  `json:"account_id"`
	TargetAccountId string  `json:"target_account_id"`
	Amount          float64 `json:"amount"`
	Reason          string  `json:"reason,omitempty"`
	Date            string  `json:"date"`
	Success         bool    `json:"success"`
}

func NewTransferHistoryOutput(id string, accountId string, targetAccountId string, amount float64, date string, isSuccessful bool, reason string) *TransferHistoryOutput {
	return &TransferHistoryOutput{
		Id:              id,
		AccountId:       accountId,
		TargetAccountId: targetAccountId,
		Reason:          reason,
		Amount:          amount,
		Date:            date,
		Success:         isSuccessful,
	}
}
