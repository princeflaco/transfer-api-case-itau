package domain

type Customer struct {
	Id        string
	Name      string
	AccountId string
}

func NewCustomer(id string, name string, accountId string) *Customer {
	return &Customer{
		Id:        id,
		Name:      name,
		AccountId: accountId,
	}
}
