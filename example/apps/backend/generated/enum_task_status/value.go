package enum_task_status

import (
	"fmt"
)

type Value string

const (
	Pending    Value = "pending"
	InProgress Value = "in_progress"
	Completed  Value = "completed"
	Cancelled  Value = "cancelled"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Pending:
		return string(v), nil
	case InProgress:
		return string(v), nil
	case Completed:
		return string(v), nil
	case Cancelled:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_task_status.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Pending:
		return nil
	case InProgress:
		return nil
	case Completed:
		return nil
	case Cancelled:
		return nil
	default:
		return fmt.Errorf("invalid enum_task_status.Value: %s", v)
	}
}
