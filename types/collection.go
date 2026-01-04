package types

import (
	"errors"
	"fmt"
)

type CollectionType string

const (
	CollectionTypeMongo  CollectionType = "mongo"
	CollectionTypeCustom CollectionType = "custom"
)

type Collection struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

func (c *Collection) Validate(registry Registry) error {
	if c.Name == "" {
		return errors.New("collection name is required")
	}

	switch CollectionType(c.Type) {
	case CollectionTypeMongo:
		return nil
	case CollectionTypeCustom:
		return nil
	default:
		return fmt.Errorf("collection type '%s' is not valid", c.Type)
	}
}
