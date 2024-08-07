package output

type GetCustomerByAccountIdOutput struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}

func NewGetCustomerByAccountIdOutput(id string, name string, accountNumber string, balance float64) *GetCustomerByAccountIdOutput {
	return &GetCustomerByAccountIdOutput{
		Id:            id,
		Name:          name,
		AccountNumber: accountNumber,
		Balance:       balance,
	}
}
