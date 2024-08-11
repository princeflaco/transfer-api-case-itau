package errors

import "fmt"

type DuplicatedEntityError struct {
	Id string
}

func NewDuplicatedEntityError(id string) *DuplicatedEntityError {
	return &DuplicatedEntityError{
		Id: id,
	}
}

func (err *DuplicatedEntityError) Error() string {
	return fmt.Sprintf("Entity already exists with this id: %s", err.Id)
}
