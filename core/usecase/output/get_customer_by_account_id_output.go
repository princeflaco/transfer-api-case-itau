package output

type GetCustomerOutput struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}

func NewGetCustomerOutput(id string, name string, accountNumber string, balance float64) *GetCustomerOutput {
	return &GetCustomerOutput{
		Id:            id,
		Name:          name,
		AccountNumber: accountNumber,
		Balance:       balance,
	}
}
