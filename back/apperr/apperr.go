package apperr

import (
	"fmt"

	"github.com/samber/lo"
)

type Code string

func (code Code) String() string {
	return string(code)
}

type Error struct {
	Code    Code
	Message string
}

func (err *Error) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("[%s]", err.Code)
	}
	return fmt.Sprintf("[%s] %s", err.Code, err.Message)
}

func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Is(err error, code Code) bool {
	aerr, ok := As(err)
	return ok && aerr.Code == code
}

func As(err error) (aerr *Error, ok bool) {
	aerr, ok = lo.ErrorsAs[*Error](err)
	return
}
