package model_template_kotlin

import (
	"strings"

	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

func getKotlinType(f types.Field, registry types.Registry) string {
	if elemType, isArray := f.ParseList(); isArray {
		return "List<" + getKotlinType(types.Field{Type: elemType}, registry) + ">"
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "Map<String, " + getKotlinType(types.Field{Type: recordType}, registry) + ">"
	}
	if _, isRef := f.ParseRef(); isRef {
		return "String"
	}
	return resolveKotlinType(f, registry)
}

func getKotlinDerefType(f types.Field, registry types.Registry) string {
	if elemType, isArray := f.ParseList(); isArray {
		return "List<" + getKotlinDerefType(types.Field{Type: elemType}, registry) + ">"
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "Map<String, " + getKotlinType(types.Field{Type: recordType}, registry) + ">"
	}
	if refFieldType, isRef := f.ParseRef(); isRef {
		return utils.UCC(refFieldType)
	}

	return resolveKotlinType(f, registry)
}

func resolveKotlinType(f types.Field, registry types.Registry) string {
	fieldTypeTest := strings.TrimPrefix(f.Type, "[]")
	switch types.FieldType(fieldTypeTest) {
	case types.FieldTypeBool:
		return "Boolean"
	case types.FieldTypeString:
		return "String"
	case types.FieldTypeInt:
		return "Int"
	case types.FieldTypeInt32:
		return "Int"
	case types.FieldTypeInt64:
		return "Long"
	case types.FieldTypeTimestamp:
		return "String" // You might want to use a proper date type like LocalDateTime
	default:
		obj, ok := registry.Get(fieldTypeTest)
		if ok {
			return obj.Name
		}
		enum, ok := registry.GetEnum(fieldTypeTest)
		if ok {
			return enum.Name
		}
		return "Any"
	}
}

func GetKotlinEnumFilename(enum types.Enum) string {
	return utils.UCC(enum.Name)
}

func tidyKotlinFile(file []byte) []byte {
	// trim leading and trailing empty lines, except for the last line
	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	var cleanedLines []string
	var foundFirst bool
	prevLine := ""
	for _, line := range lines {
		if line == "" && !foundFirst {
			continue
		}
		foundFirst = true
		if line == "" && prevLine == "" {
			continue
		}
		cleanedLines = append(cleanedLines, line)
		prevLine = line
	}
	return []byte(strings.Join(cleanedLines, "\n"))
}
