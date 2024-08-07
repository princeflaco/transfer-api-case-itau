package output

import "encoding/json"

type CreateCustomerOutput struct {
	Id        string `json:"id"`
	AccountId string `json:"account_id"`
}

func NewCreateCustomerOutput(id string, accountId string) *CreateCustomerOutput {
	return &CreateCustomerOutput{
		Id:        id,
		AccountId: accountId,
	}
}

func (o *CreateCustomerOutput) ToBytes() ([]byte, error) {
	return json.Marshal(o)
}
