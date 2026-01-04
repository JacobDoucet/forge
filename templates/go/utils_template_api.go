package model_template_go

import (
	"strings"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

func fieldToApiModelTransform(field types.Field, registry types.Registry) string {
	fieldName := "r." + templates.TrimIDSuffix(utils.UCC(field.Name))
	if _, isList := field.ParseList(); isList {
		return `val := make(` + getGoDerefType(field, registry, "Model") + `, 0)
		var err error
		for _, rr := range *` + fieldName + ` {
			nextVal, nextErr := rr.ToModel()
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}`
	}
	return strings.Join([]string{
		"val, toModelErr := " + fieldName + ".ToModel()",
		"if toModelErr != nil {",
		"err = errors.Join(err, toModelErr)",
		"}",
	}, "\n")
}

func fieldToHTTPApiModelTransform(field types.Field, registry types.Registry) string {
	fieldName := "r." + templates.TrimIDSuffix(utils.UCC(field.Name))
	if _, isList := field.ParseList(); isList {
		return `val := make(` + getGoDerefType(field, registry, "HTTPRecord") + `, 0)
		for _, rr := range *` + fieldName + ` {
			nextVal, nextErr := rr.ToHTTPRecord(refProjection)
			if nextErr != nil {
				err = errors.Join(err, nextErr)
			}
			val = append(val, nextVal)
		}`
	}
	return strings.Join([]string{
		"val, toHTTPRecordErr := " + fieldName + ".ToHTTPRecord(refProjection)",
		"if toHTTPRecordErr != nil {",
		"err = errors.Join(err, toHTTPRecordErr)",
		"}",
	}, "\n")
}
