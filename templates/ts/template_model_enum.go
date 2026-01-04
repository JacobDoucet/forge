package model_template_ts

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed model/enum.ts.tmpl
var modelEnumTSTemplate string

func NewModelEnumTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__obj_api_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetObjModelFilename": GetTSModelFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getTSDerefType(field, registry)
			},
			"EnumFieldName": func(e types.Enum, fieldName string) string {
				if e.Type == types.FieldTypeInt {
					nameSplit := strings.Split(string(fieldName), "=")
					return strings.TrimSpace(nameSplit[0])
				}
				return fieldName
			},
			"FmtEnumValueType": func(e types.Enum) string {
				var t []string
				for _, v := range e.Values {
					t = append(t, fmtEnumValue(e, v))
				}
				return strings.Join(t, "\n    | ")
			},
			"FmtEnumValue": func(e types.Enum, v string) string {
				return fmtEnumValue(e, v)
			},
			"FmtEnumValueName": func(e types.Enum, v string) string {
				return utils.UCC(e.Name) + utils.UCC(fmtEnumValueName(e, v))
			},
		})).
		Parse(modelEnumTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}

func fmtEnumValue(e types.Enum, v string) string {
	if e.Type == types.FieldTypeString {
		return "'" + v + "'"
	}
	if e.Type == types.FieldTypeInt {
		nameSplit := strings.Split(v, "=")
		return strings.TrimSpace(nameSplit[len(nameSplit)-1])
	}
	return v
}

func fmtEnumValueName(e types.Enum, v string) string {
	if e.Type == types.FieldTypeInt {
		nameSplit := strings.Split(v, "=")
		return strings.TrimSpace(nameSplit[0])
	}
	return v
}
