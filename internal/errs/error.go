package errs

import (
	"fmt"

	"github.com/Chihaya-Anon123/task_manager/internal/code"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInvalidParams  = New(code.CodeInvalidParams, "invalid params")
	ErrUnauthorized   = New(code.CodeUnauthorized, "unauthorized")
	ErrNotFound       = New(code.CodeNotFound, "resource not found")
	ErrInternalServer = New(code.CodeInternalServer, "internal server error")
	ErrDBError        = New(code.CodeDBError, "database error")
)
