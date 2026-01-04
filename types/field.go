package types

import (
	"errors"
	"fmt"
	"regexp"
)

type Field struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	Permissions ObjectPermissions `yaml:"permissions"`
	Immutable   bool              `yaml:"immutable"`
	Required    bool              `yaml:"required"`
	Values      []string          `yaml:"values"` // Only applicable to Enums
}

type FieldType string

const (
	FieldTypeBool      FieldType = "bool"
	FieldTypeString    FieldType = "string"
	FieldTypeInt       FieldType = "int"
	FieldTypeInt32     FieldType = "int32"
	FieldTypeInt64     FieldType = "int64"
	FieldTypeTimestamp FieldType = "timestamp"
	FieldActorRole     FieldType = "permissions.ActorRole"
)

type RootFieldType string

const (
	RootFieldTypePrimitive RootFieldType = "primitive"
	RootFieldTypeObject    RootFieldType = "object"
	RootFieldTypeEnum      RootFieldType = "enum"
)

func (f Field) ResolveRootType(registry Registry) (string, RootFieldType, error) {
	var err error
	if f.Name == "" {
		err = errors.Join(err, fmt.Errorf("field name is required"))
	}
	if f.Type == "" {
		err = errors.Join(err, fmt.Errorf("field type is required"))
	}

	elemType, isArray := f.ParseList()
	if isArray {
		rootType, rootFieldType, arrErr := Field{Name: f.Name, Type: elemType}.ResolveRootType(registry)
		if arrErr != nil {
			err = errors.Join(err, arrErr)
		}
		return rootType, rootFieldType, err
	}

	valType, isKeyVal := f.ParseKeyVal()
	if isKeyVal {
		rootType, rootFieldType, keyValErr := Field{Name: f.Name, Type: valType}.ResolveRootType(registry)
		if keyValErr != nil {
			err = errors.Join(err, keyValErr)
		}
		return rootType, rootFieldType, err
	}

	refType, isRef := f.ParseRef()
	if isRef {
		_, ok := registry.Get(refType)
		if !ok {
			err = errors.Join(err, fmt.Errorf("field type %s is not a valid reference", f.Type))
		}
		return refType, RootFieldTypeObject, err
	}

	if f.IsPrimitive() {
		return f.Type, RootFieldTypePrimitive, nil
	}
	_, ok := registry.Get(f.Type)
	if ok {
		return f.Type, RootFieldTypeObject, nil
	}
	_, ok = registry.GetEnum(f.Type)
	if ok {
		return f.Type, RootFieldTypeEnum, nil
	}
	return "", RootFieldTypePrimitive, fmt.Errorf("field type %s is not a primitive or a registered object", f.Type)
}

func (f Field) Validate(registry Registry) error {
	_, _, err := f.ResolveRootType(registry)
	return err
}

func (f Field) validateType(registry Registry) error {
	if f.IsPrimitive() {
		return nil
	}
	_, ok := registry.Get(f.Type)
	if ok {
		return nil
	}
	return fmt.Errorf("field type %s is not a primitive or a registered object", f.Type)
}

func (f Field) IsPrimitive() bool {
	switch FieldType(f.Type) {
	case FieldTypeBool, FieldTypeString, FieldTypeInt, FieldTypeInt32, FieldTypeInt64, FieldTypeTimestamp:
		return true
	default:
		return false
	}
}

func (f Field) IsObject(registry Registry) bool {
	_, ok := registry.Get(f.Type)
	return ok
}

func (f Field) IsEnum(registry Registry) bool {
	_, ok := registry.GetEnum(f.Type)
	return ok
}

func (f Field) IsRef() bool {
	pattern, _ := regexp.Compile("^(Parent)*Ref<(.+)>$")
	match := pattern.FindStringSubmatch(f.Type)
	return len(match) > 1
}

func (f Field) ParseEnum() (string, bool) {
	pattern, _ := regexp.Compile("^Enum<(.+)>$")
	match := pattern.FindStringSubmatch(f.Type)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

func (f Field) ParseList() (string, bool) {
	pattern, _ := regexp.Compile("^List<(.+)>$")
	match := pattern.FindStringSubmatch(f.Type)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

func (f Field) ParseKeyVal() (string, bool) {
	pattern, _ := regexp.Compile("^KeyVal<(.+)>$")
	match := pattern.FindStringSubmatch(f.Type)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

func (f Field) ParseRef() (string, bool) {
	field := f

	if elemType, isList := f.ParseList(); isList {
		return Field{Type: elemType}.ParseRef()
	}
	if recordType, isKeyVal := field.ParseKeyVal(); isKeyVal {
		return Field{Type: recordType}.ParseRef()
	}

	pattern, _ := regexp.Compile("^(Parent)*Ref<(.+)>$")
	match := pattern.FindStringSubmatch(field.Type)
	if len(match) > 1 {
		return match[len(match)-1], true
	}
	return "", false
}

func (f Field) ParseParentRef() (string, bool) {
	pattern, _ := regexp.Compile("^ParentRef<(.+)>$")
	match := pattern.FindStringSubmatch(f.Type)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}
