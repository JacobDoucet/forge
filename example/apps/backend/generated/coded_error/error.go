package coded_error

import (
	"errors"
	"fmt"
	"strings"
)

type Code string

const (
	CodeInvalidRequest      Code = "INVALID_REQUEST"
	CodeUnauthorized        Code = "UNAUTHORIZED"
	CodeNotFound            Code = "NOT_FOUND"
	CodeMethodNotAllowed    Code = "METHOD_NOT_ALLOWED"
	CodeEntityAlreadyExists Code = "ENTITY_ALREADY_EXISTS"
	CodeUnexpected          Code = "UNEXPECTED"
	CodeProjectNotFound     Code = "PROJECT_NOT_FOUND"
	CodeInvalidTaskStatus   Code = "INVALID_TASK_STATUS"
	CodeTaskNotFound        Code = "TASK_NOT_FOUND"
)

type customError struct {
	errorCode  Code
	httpStatus int
	message    string
}

// Error implements the error interface.
func (e *customError) Error() string {
	return fmt.Sprintf("errorCode=%s, httpStatus=%d, message=%s", e.errorCode, e.httpStatus, e.message)
}

func newCustomError(errorCode Code, httpStatus int, message string) error {
	return &customError{
		errorCode:  errorCode,
		httpStatus: httpStatus,
		message:    message,
	}
}

// ResolveHTTPStatus extracts the HTTP status from an error, even if it's wrapped.
func ResolveHTTPStatus(err error) int {
	var customErr *customError
	if errors.As(err, &customErr) {
		return customErr.httpStatus
	}
	return 500
}

// ResolveErrorCode extracts the error code from an error, even if it's wrapped.
func ResolveErrorCode(err error) Code {
	if err == nil {
		return ""
	}
	var customErr *customError
	if errors.As(err, &customErr) {
		return customErr.errorCode
	}
	return CodeUnexpected
}

func ResolveErrorCodeAsString(err error) string {
	return string(ResolveErrorCode(err))
}

func NewInvalidRequestError(message ...string) error {
	msg := "The request is invalid"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeInvalidRequest, 400, msg)
}

func NewUnauthorizedError(message ...string) error {
	msg := "The request is unauthorized"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeUnauthorized, 401, msg)
}

func NewNotFoundError(message ...string) error {
	msg := "The requested resource was not found"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeNotFound, 404, msg)
}

func NewMethodNotAllowedError(message ...string) error {
	msg := "The method is not allowed"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeMethodNotAllowed, 405, msg)
}

func NewEntityAlreadyExistsError(message ...string) error {
	msg := "The entity already exists"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeEntityAlreadyExists, 409, msg)
}

func NewProjectNotFoundError(message ...string) error {
	msg := "The requested project was not found"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeProjectNotFound, 500, msg)
}

func NewUnexpectedError(message ...string) error {
	msg := "An unexpected error occurred"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeUnexpected, 500, msg)
}

func NewTaskNotFoundError(message ...string) error {
	msg := "The requested task was not found"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeTaskNotFound, 500, msg)
}

func NewInvalidTaskStatusError(message ...string) error {
	msg := "The provided task status is invalid"
	if len(message) > 0 {
		msg = msg + ": " + strings.Join(message, ". ")
	}
	return newCustomError(CodeInvalidTaskStatus, 500, msg)
}
