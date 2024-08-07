package output

type TransferHistoryOutput struct {
	Id              string  `json:"id"`
	TargetAccountId string  `json:"target_account_id"`
	Amount          float64 `json:"amount"`
	Date            string  `json:"date"`
	Success         bool    `json:"success"`
}

func NewTransferHistoryOutput(id string, targetAccountId string, amount float64, date string, isSuccessful bool) *TransferHistoryOutput {
	return &TransferHistoryOutput{
		Id:              id,
		TargetAccountId: targetAccountId,
		Amount:          amount,
		Date:            date,
		Success:         isSuccessful,
	}
}
