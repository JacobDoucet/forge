package types

import (
	"errors"
	"fmt"
)

type EnumRegistry interface {
	Get(name string) (Enum, bool)
	Register(enum Enum) error
	List() []Enum
}

func NewEnumRegistry() EnumRegistry {
	return &enumRegistry{
		enums: make(map[string]Enum),
	}
}

type enumRegistry struct {
	enums map[string]Enum
}

func (r *enumRegistry) Get(name string) (Enum, bool) {
	enum, ok := r.enums[name]
	return enum, ok
}

func (r *enumRegistry) Register(enum Enum) error {
	if _, ok := r.enums[enum.Name]; ok {
		return fmt.Errorf("enum %s already registered", enum.Name)
	}
	if err := enum.Validate(); err != nil {
		return errors.Join(fmt.Errorf("enum %s is invalid", enum.Name), err)
	}
	r.enums[enum.Name] = enum
	return nil
}

func (r *enumRegistry) List() []Enum {
	var enums []Enum
	for _, enum := range r.enums {
		enums = append(enums, enum)
	}
	return enums
}
