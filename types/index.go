package types

import (
	"fmt"
	"strings"

	"github.com/JacobDoucet/forge/utils"
)

type Index struct {
	Name        string       `yaml:"name"`
	Unique      bool         `yaml:"unique"`
	Expiration  bool         `yaml:"expiration"`
	Fields      []IndexField `yaml:"fields"`
	IncludeNull bool         `yaml:"includeNull"`
}

type IndexField struct {
	Name  string `yaml:"name"`
	Order int    `yaml:"order"`
}

func (f *IndexField) FormatSortParam() string {
	return utils.UCC(strings.ReplaceAll(f.Name, ".", "_"))
}

func (i *Index) Validate(registry Registry, obj *Object) error {
	if i.Name == "" {
		return fmt.Errorf("index name is required")
	}
	if len(i.Fields) == 0 {
		return fmt.Errorf("index fields are required")
	}
	for iField, field := range i.Fields {
		if err := validateIndexField(strings.Split(field.Name, "."), registry, *obj); err != nil {
			return fmt.Errorf("index field %s is invalid: %w", field.Name, err)
		}
		if field.Order < 0 {
			field.Order = -1
		} else {
			field.Order = 1
		}
		i.Fields[iField] = field
	}
	return nil
}

func validateIndexField(namePath []string, registry Registry, obj Object) error {
	if len(namePath) == 0 {
		return fmt.Errorf("index field name is required")
	}
	fieldName := namePath[0]
	if len(namePath) == 1 {
		if _, ok := obj.fieldMap[fieldName]; !ok {
			return fmt.Errorf("field %s.%s is not defined", obj.Name, fieldName)
		}
		return nil
	}

	field, ok := obj.GetField(fieldName)
	if !ok {
		return fmt.Errorf("field %s.%s is not defined", obj.Name, fieldName)
	}

	rootType, _, _ := field.ResolveRootType(registry)
	nextObj, ok := registry.Get(rootType)
	if !ok {
		return fmt.Errorf("field %s.%s of type %s is not a valid reference", obj.Name, fieldName, field.Type)
	}

	return validateIndexField(namePath[1:], registry, nextObj)
}
