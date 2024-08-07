package input

import (
	"encoding/json"
	"transfer-api/core/errors"
)

type CreateCustomerInput struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
	Balance   int    `json:"balance"`
}

func (i CreateCustomerInput) Validate() error {
	if i.Name == "" {
		return errors.NewInvalidFieldError("name", "Should not be empty")
	}
	return nil
}

func (i CreateCustomerInput) ToBytes() ([]byte, error) {
	return json.Marshal(i)
}
