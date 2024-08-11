package input

import "transfer-api/core/errors"

type TransferInput struct {
	TargetAccountId string  `json:"target_account_id" required:"true"`
	Amount          float64 `json:"amount" required:"true"`
}

func (input *TransferInput) Validate() []errors.InvalidFieldError {
	var fieldErrors []errors.InvalidFieldError
	if input.TargetAccountId == "" {
		fieldErrors = append(fieldErrors, *errors.NewInvalidFieldError("numero_conta_destino", "Should not be empty or nil"))
	}
	if input.Amount <= 0 {
		fieldErrors = append(fieldErrors, *errors.NewInvalidFieldError("valor", "Should be positive"))
	}
	return fieldErrors
}
