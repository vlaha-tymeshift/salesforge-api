package errors

import (
	"fmt"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	// Check if the nested error is an AppError.
	// If it is, include only the innermost error message.
	// If it's not, include the error message as usual.
	if nestedErr, ok := e.Err.(*AppError); ok {
		return fmt.Sprintf("code: %d, message: %s, error: %s", e.Code, e.Message, nestedErr.Err.Error())
	}
	return fmt.Sprintf("code: %d, message: %s, error: %s", e.Code, e.Message, e.Err.Error())
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Unwrap() error {
	return e.Err
}
