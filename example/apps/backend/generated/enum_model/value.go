package enum_model

import (
	"fmt"
)

type Value string

const (
	Project Value = "Project"
	Task    Value = "Task"
	User    Value = "User"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Project:
		return string(v), nil
	case Task:
		return string(v), nil
	case User:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Project:
		return nil
	case Task:
		return nil
	case User:
		return nil
	default:
		return fmt.Errorf("invalid enum_model.Value: %s", v)
	}
}
