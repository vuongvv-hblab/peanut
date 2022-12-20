package apierrors

import (
	"errors"
	"fmt"
)

type ApiError struct {
	errType ErrorType
	err     error
}

// Unwrap implements the errors.Wrapper interface.
func (e *ApiError) Unwrap() error {
	return e.err
}

// Error implements the error interface.
func (e *ApiError) Error() string {
	msg := fmt.Sprintf("%v %s ", e.errType.HTTPCode(), e.errType.Code())
	if e.err != nil {
		msg += e.err.Error()
	}

	return msg
}

func New(errType ErrorType, err error) *ApiError {
	return &ApiError{
		errType: errType,
		err:     err,
	}
}

func NewErrorf(errType ErrorType, format string, a ...interface{}) *ApiError {
	return New(errType, fmt.Errorf(format, a...))
}

func ErrType(err error) ErrorType {
	var apiError *ApiError
	if errors.As(err, &apiError) {
		return apiError.errType
	}
	return InternalError
}

func IsErrType(err *ApiError, errType ErrorType) bool {
	return err.errType == errType
}
