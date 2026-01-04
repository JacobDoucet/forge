package enum_event_type

import (
	"fmt"
)

type Value string

const (
	ProjectCreated    Value = "ProjectCreated"
	ProjectUpdated    Value = "ProjectUpdated"
	TaskCreated       Value = "TaskCreated"
	TaskDeleted       Value = "TaskDeleted"
	TaskStatusChanged Value = "TaskStatusChanged"
	TaskUpdated       Value = "TaskUpdated"
)

func (v Value) ToString() (string, error) {
	switch v {
	case ProjectCreated:
		return string(v), nil
	case ProjectUpdated:
		return string(v), nil
	case TaskCreated:
		return string(v), nil
	case TaskDeleted:
		return string(v), nil
	case TaskStatusChanged:
		return string(v), nil
	case TaskUpdated:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_event_type.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case ProjectCreated:
		return nil
	case ProjectUpdated:
		return nil
	case TaskCreated:
		return nil
	case TaskDeleted:
		return nil
	case TaskStatusChanged:
		return nil
	case TaskUpdated:
		return nil
	default:
		return fmt.Errorf("invalid enum_event_type.Value: %s", v)
	}
}
