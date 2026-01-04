package enum_role

import (
	"fmt"
)

type Value string

const (
	Super Value = "Super"
	Admin Value = "Admin"
	User  Value = "User"
	Guest Value = "Guest"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Super:
		return string(v), nil
	case Admin:
		return string(v), nil
	case User:
		return string(v), nil
	case Guest:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_role.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Super:
		return nil
	case Admin:
		return nil
	case User:
		return nil
	case Guest:
		return nil
	default:
		return fmt.Errorf("invalid enum_role.Value: %s", v)
	}
}
