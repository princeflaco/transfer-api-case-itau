package errors

type NotFoundError struct {
	Entity string
	Id     string
}

var _ error = (*NotFoundError)(nil)

func NewNotFoundError(entity string, id string) *NotFoundError {
	return &NotFoundError{
		Entity: entity,
		Id:     id,
	}
}

func (err *NotFoundError) Error() string {
	return "Entity not found: " + err.Entity + ":" + err.Id
}
