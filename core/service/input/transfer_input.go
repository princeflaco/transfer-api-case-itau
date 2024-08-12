package input

import "transfer-api/core/errors"

type TransferInput struct {
	AccountId       string  `json:"account_id" required:"true"`
	TargetAccountId string  `json:"target_account_id" required:"true"`
	Amount          float64 `json:"amount" required:"true"`
}

func (input *TransferInput) Validate() []errors.InvalidFieldError {
	var fieldErrors []errors.InvalidFieldError
	if input.AccountId == "" {
		fieldErrors = append(fieldErrors, *errors.NewInvalidFieldError("account_id", "Should not be empty or nil"))
	}
	if input.TargetAccountId == "" {
		fieldErrors = append(fieldErrors, *errors.NewInvalidFieldError("target_account_id", "Should not be empty or nil"))
	}
	if input.Amount <= 0 {
		fieldErrors = append(fieldErrors, *errors.NewInvalidFieldError("amount", "Should be positive"))
	}
	return fieldErrors
}
