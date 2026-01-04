package types

import (
	"fmt"
	"sort"
)

type ErrorRegistry interface {
	Get(name string) (CustomError, bool)
	Register(error CustomError) error
	List() []CustomError
}

func NewErrorRegistry() ErrorRegistry {
	r := errorRegistry{}
	_ = r.Register(CustomError{
		Code:       "UNEXPECTED",
		HttpStatus: 500,
		Message:    "An unexpected error occurred",
	})
	_ = r.Register(CustomError{
		Code:       "NOT_FOUND",
		HttpStatus: 404,
		Message:    "The requested resource was not found",
	})
	_ = r.Register(CustomError{
		Code:       "INVALID_REQUEST",
		HttpStatus: 400,
		Message:    "The request is invalid",
	})
	_ = r.Register(CustomError{
		Code:       "ENTITY_ALREADY_EXISTS",
		HttpStatus: 409,
		Message:    "The entity already exists",
	})
	_ = r.Register(CustomError{
		Code:       "METHOD_NOT_ALLOWED",
		HttpStatus: 405,
		Message:    "The method is not allowed",
	})
	_ = r.Register(CustomError{
		Code:       "UNAUTHORIZED",
		HttpStatus: 401,
		Message:    "The request is unauthorized",
	})

	return r
}

type errorRegistry map[string]CustomError

func (r errorRegistry) Get(name string) (CustomError, bool) {
	e, ok := r[name]
	return e, ok
}

func (r errorRegistry) Register(error CustomError) error {
	if _, ok := r[error.Code]; ok {
		return fmt.Errorf("error %s already registered", error.Code)
	}
	if err := error.Validate(); err != nil {
		return fmt.Errorf("error %s is invalid: %w", error.Code, err)
	}
	r[error.Code] = error
	return nil
}

func (r errorRegistry) List() []CustomError {
	var errors []CustomError
	for _, e := range r {
		errors = append(errors, e)
	}
	sort.Slice(errors, func(i, j int) bool {
		return errors[i].HttpStatus < errors[j].HttpStatus
	})
	return errors
}
