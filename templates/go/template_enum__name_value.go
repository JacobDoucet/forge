package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed enum__name/value.go.tmpl
var enumNameValueTemplate string

func NewEnumNameValueGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("enum__value").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						packageFmt,
					}
				},
				"EnumFieldName": func(e types.Enum, fieldName string) string {
					if e.Type == types.FieldTypeInt {
						nameSplit := strings.Split(string(fieldName), "=")
						return strings.TrimSpace(nameSplit[0])
					}
					return fieldName
				},
				"EnumType": func(e types.Enum) string {
					return string(e.Type)
				},
				"EnumStringFmt": func(e types.Enum) string {
					if e.Type == types.FieldTypeString {
						return "%s"
					}
					if e.Type == types.FieldTypeInt {
						return "%d"
					}
					return "%v"
				},
				"FmtEnumValue": func(e types.Enum, v string) string {
					if e.Type == types.FieldTypeString {
						return "\"" + v + "\""
					}
					if e.Type == types.FieldTypeInt {
						nameSplit := strings.Split(v, "=")
						return strings.TrimSpace(nameSplit[len(nameSplit)-1])
					}
					return v
				},
				"EnumDefaultValue": func(e types.Enum) string {
					if e.Type == types.FieldTypeString {
						return "\"\""
					}
					return "0"
				},
				"ListObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasHTTPMethods() {
							objs = append(objs, obj)
						}
					}
					return objs
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(enumNameValueTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = apiTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}
