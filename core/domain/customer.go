package domain

import (
	"github.com/google/uuid"
)

type Customer struct {
	Id   string
	Name string
}

func NewCustomer(id string, name string) *Customer {
	if id == "" {
		id = uuid.NewString()
	}
	return &Customer{
		Id:   id,
		Name: name,
	}
}
