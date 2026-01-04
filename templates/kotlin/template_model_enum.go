package model_template_kotlin

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed model/enum.kt.tmpl
var modelEnumKotlinTemplate string

func NewModelEnumKotlinGenerator(registry types.Registry) (templates.KotlinGenFunc, error) {
	tmpl, err := template.
		New("package__enum_kotlin").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPkgName": func() string {
				return registry.GetKotlinPkgRoot() + ".model"
			},
			"GetObjModelFilename": GetKotlinModelFilename,
			"GetFieldType": func(field types.Field) string {
				return getKotlinType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getKotlinDerefType(field, registry)
			},
			"GetEnumValueType": func(e types.Enum) string {
				if e.Type == types.FieldTypeString {
					return "String"
				}
				return "Int"
			},
			"FmtEnumValue": func(e types.Enum, v string) string {
				return fmtKotlinEnumValue(e, v)
			},
			"FmtEnumValueName": func(e types.Enum, v string) string {
				if e.Type == types.FieldTypeInt {
					nameSplit := strings.Split(string(v), "=")
					return utils.UCC(strings.TrimSpace(nameSplit[0]))
				}
				return utils.UCC(v)
			},
		})).
		Parse(modelEnumKotlinTemplate)

	return func(ctx templates.KotlinTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyKotlinFile(buf.Bytes()), err
	}, err
}

func fmtKotlinEnumValue(e types.Enum, v string) string {
	if e.Type == types.FieldTypeString {
		return "\"" + v + "\""
	}
	if e.Type == types.FieldTypeInt {
		nameSplit := strings.Split(v, "=")
		return strings.TrimSpace(nameSplit[len(nameSplit)-1])
	}
	return v
}
