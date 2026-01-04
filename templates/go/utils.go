package model_template_go

import (
	"fmt"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

func getGoDerefType(f types.Field, registry types.Registry, objectType string) string {
	if elemType, isList := f.ParseList(); isList {
		return "[]" + getGoDerefType(types.Field{Type: elemType}, registry, objectType)
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "map[string]" + getGoDerefType(types.Field{Type: recordType}, registry, objectType)
	}
	if refFieldType, isRef := f.ParseRef(); isRef {
		return utils.SC(refFieldType) + "." + utils.UCC(objectType)
	}

	return resolveGoType(f, registry, objectType)
}

func getGoModelType(f types.Field, registry types.Registry) string {
	if elemType, isList := f.ParseList(); isList {
		return "[]" + getGoModelType(types.Field{Type: elemType}, registry)
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "map[string]" + getGoModelType(types.Field{Type: recordType}, registry)
	}
	if enumType, isEnum := f.ParseEnum(); isEnum {
		return utils.SC(enumType) + ".Value"
	}
	if _, isRef := f.ParseRef(); isRef {
		return "string"
	}

	return resolveGoType(f, registry, "Model")
}

func getGoHTTPType(f types.Field, registry types.Registry) string {
	if elemType, isList := f.ParseList(); isList {
		elemField := types.Field{Type: elemType}
		return "[]" + getGoHTTPType(elemField, registry)
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "map[string]" + getGoHTTPType(types.Field{Type: recordType}, registry)
	}
	if _, isRef := f.ParseRef(); isRef {
		return "string"
	}

	return resolveGoType(f, registry, "HTTPRecord")
}

func getGoMongoType(f types.Field, registry types.Registry) string {
	if elemType, isList := f.ParseList(); isList {
		return "[]" + getGoMongoType(types.Field{Type: elemType}, registry)
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "map[string]" + getGoMongoType(types.Field{Type: recordType}, registry)
	}
	if _, isRef := f.ParseRef(); isRef {
		return "primitive.ObjectID"
	}

	return resolveGoType(f, registry, "MongoRecord")
}

func resolveGoType(f types.Field, registry types.Registry, objectType string) string {
	switch types.FieldType(f.Type) {
	case types.FieldTypeBool:
		return "bool"
	case types.FieldTypeString:
		return "string"
	case types.FieldTypeInt:
		return "int"
	case types.FieldTypeInt32:
		return "int32"
	case types.FieldTypeInt64:
		return "int64"
	case types.FieldTypeTimestamp:
		return "time.Time"
	default:
		obj, ok := registry.Get(f.Type)
		if ok {
			return templates.GetModelPackageName(obj) + "." + objectType
		}
		enum, ok := registry.GetEnum(f.Type)
		if ok {
			return templates.GetEnumPackageName(enum) + ".Value"
		}
		fmt.Println("Field type not found in registry: ", f.Name, f.Type)
		return "interface{}"
	}
}
