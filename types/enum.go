package types

import (
	"errors"
	"fmt"

	"d3tech.com/platform/utils"
)

type Enum struct {
	Name   string    `yaml:"name"`
	Type   FieldType `yaml:"type"`
	Values []string  `yaml:"values"`
}

func (e *Enum) GetPkgName() string {
	return fmt.Sprintf("enum_%s", utils.SC(e.Name))
}

func (e *Enum) Validate() error {
	if e.Name == "" {
		return errors.New("missing name")
	}
	if e.Type == "" {
		return errors.New("missing type")
	}
	if e.Type != FieldTypeString && e.Type != FieldTypeInt && e.Type != FieldTypeInt32 && e.Type != FieldTypeInt64 {
		return fmt.Errorf("invalid type %s, only %s, %s, %s or %s allowed", e.Type, FieldTypeString, FieldTypeInt, FieldTypeInt32, FieldTypeInt64)
	}
	if len(e.Values) == 0 {
		return errors.New("missing values")
	}
	return nil
}
