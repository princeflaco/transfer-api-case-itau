package errors

import "fmt"

type InvalidFieldError struct {
	FieldName string
	AsIs      string
}

var _ error = (*InvalidFieldError)(nil)

func NewInvalidFieldError(fieldName, asIs string) *InvalidFieldError {
	return &InvalidFieldError{
		FieldName: fieldName,
		AsIs:      asIs,
	}
}

func (e *InvalidFieldError) Error() string {
	return fmt.Sprintf("Invalid field `%s`, as is: %s", e.FieldName, e.AsIs)
}
