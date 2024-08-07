package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id   string
	Name string
}

func NewUser(id string, name string) *User {
	if id == "" {
		id = uuid.NewString()
	}
	return &User{
		Id:   id,
		Name: name,
	}
}
