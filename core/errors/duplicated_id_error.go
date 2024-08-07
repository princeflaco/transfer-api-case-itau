package errors

type DuplicatedIdError struct {
	Id string
}

func NewDuplicatedIdError(id string) *DuplicatedIdError {
	return &DuplicatedIdError{
		Id: id,
	}
}

func (err *DuplicatedIdError) Error() string {
	return err.Id
}
