package output

import "encoding/json"

type CreateUserOutput struct {
	Id        string `json:"id"`
	AccountId string `json:"account_id"`
}

func NewCreateUserOutput(id string, accountId string) *CreateUserOutput {
	return &CreateUserOutput{
		Id:        id,
		AccountId: accountId,
	}
}

func (o *CreateUserOutput) ToBytes() ([]byte, error) {
	return json.Marshal(o)
}
