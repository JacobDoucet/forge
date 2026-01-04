package enum_task_priority

import (
	"fmt"
)

type Value string

const (
	Low    Value = "low"
	Medium Value = "medium"
	High   Value = "high"
	Urgent Value = "urgent"
)

func (v Value) ToString() (string, error) {
	switch v {
	case Low:
		return string(v), nil
	case Medium:
		return string(v), nil
	case High:
		return string(v), nil
	case Urgent:
		return string(v), nil
	default:
		return "", fmt.Errorf("invalid enum_task_priority.Value: %s", v)
	}
}

func Validate(v Value) error {
	switch v {
	case Low:
		return nil
	case Medium:
		return nil
	case High:
		return nil
	case Urgent:
		return nil
	default:
		return fmt.Errorf("invalid enum_task_priority.Value: %s", v)
	}
}
