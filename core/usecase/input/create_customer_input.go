package input

import (
	"encoding/json"
	"transfer-api/core/errors"
)

type CreateCustomerInput struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	AccountId string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

func (i CreateCustomerInput) Validate() error {
	if i.Name == "" {
		return errors.NewInvalidFieldError("name", "Should not be empty")
	}
	if i.AccountId == "" {
		return errors.NewInvalidFieldError("account_id", "Should not be empty")
	}
	if i.Id == "" {
		return errors.NewInvalidFieldError("id", "Should not be empty")
	}
	return nil
}

func (i CreateCustomerInput) ToBytes() ([]byte, error) {
	return json.Marshal(i)
}
