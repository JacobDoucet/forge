package model_template_ts

import (
	"strings"

	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

func getTSType(f types.Field, registry types.Registry) string {
	if elemType, isArray := f.ParseList(); isArray {
		return getTSType(types.Field{Type: elemType}, registry) + "[]"
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "Record<string," + getTSType(types.Field{Type: recordType}, registry) + ">"
	}
	if _, isRef := f.ParseRef(); isRef {
		return "string"
	}
	return resolveTSType(f, registry)
}

func getTSDerefType(f types.Field, registry types.Registry) string {
	if elemType, isArray := f.ParseList(); isArray {
		return getTSDerefType(types.Field{Type: elemType}, registry) + "[]"
	}
	if recordType, isKeyVal := f.ParseKeyVal(); isKeyVal {
		return "Record<string," + getTSType(types.Field{Type: recordType}, registry) + ">"
	}
	if refFieldType, isRef := f.ParseRef(); isRef {
		return utils.UCC(refFieldType)
	}

	return resolveTSType(f, registry)
}

func resolveTSType(f types.Field, registry types.Registry) string {
	fieldTypeTest := strings.TrimPrefix(f.Type, "[]")
	switch types.FieldType(fieldTypeTest) {
	case types.FieldTypeBool:
		return "boolean"
	case types.FieldTypeString:
		return "string"
	case types.FieldTypeInt:
		return "number"
	case types.FieldTypeInt32:
		return "number"
	case types.FieldTypeInt64:
		return "number"
	case types.FieldTypeTimestamp:
		return "string"
	default:
		obj, ok := registry.Get(fieldTypeTest)
		if ok {
			return obj.Name
		}
		enum, ok := registry.GetEnum(fieldTypeTest)
		if ok {
			return enum.Name
		}
		return "any"
	}
}

func tidyTSFile(file []byte) []byte {
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
