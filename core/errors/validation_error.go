package errors

import "fmt"

type ValidationError struct {
	InvalidFields []InvalidFieldError
}

func NewValidationError(invalidFields ...InvalidFieldError) *ValidationError {
	return &ValidationError{
		InvalidFields: invalidFields,
	}
}

func (err *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %v", err.InvalidFields)
}
