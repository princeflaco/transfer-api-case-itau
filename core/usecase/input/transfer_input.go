package input

import "transfer-api/core/errors"

type TransferInput struct {
	TargetAccountId string  `json:"target_account_id"`
	Amount          float64 `json:"amount"`
}

func (input *TransferInput) Validate() error {
	if input.TargetAccountId == "" {
		return errors.NewInvalidFieldError("numero_conta_destino", "Should not be empty or nil")
	}
	if input.Amount <= 0 {
		return errors.NewInvalidFieldError("valor", "Should be positive")
	}
	return nil
}
