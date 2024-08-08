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

func (i CreateCustomerInput) Validate() []errors.InvalidFieldError {
	var errs []errors.InvalidFieldError
	if i.Name == "" {
		errs = append(errs, *errors.NewInvalidFieldError("name", "Should not be empty"))
	}
	if i.AccountId == "" {
		errs = append(errs, *errors.NewInvalidFieldError("account_id", "Should not be empty"))
	}
	if i.Id == "" {
		errs = append(errs, *errors.NewInvalidFieldError("id", "Should not be empty"))
	}
	return errs
}

func (i CreateCustomerInput) ToBytes() ([]byte, error) {
	return json.Marshal(i)
}
