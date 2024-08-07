package input

import (
	"encoding/json"
	"transfer-api/src/core/errors"
)

type CreateUserInput struct {
	Id            string `json:"id"`
	Name          string `json:"nome"`
	AccountNumber string `json:"numero_conta"`
	Balance       int    `json:"saldo"`
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
