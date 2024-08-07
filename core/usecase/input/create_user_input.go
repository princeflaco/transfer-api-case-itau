package input

import (
	"encoding/json"
	"transfer-api/core/errors"
)

type CreateUserInput struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
	Balance   int    `json:"balance"`
}

func (i CreateUserInput) Validate() error {
	if i.Name == "" {
		return errors.NewInvalidFieldError("name", "Should not be empty")
	}
	return nil
}

func (i CreateUserInput) ToBytes() ([]byte, error) {
	return json.Marshal(i)
}
