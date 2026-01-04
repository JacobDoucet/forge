package model_template_ts

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed api/obj__endpoints.ts.tmpl
var apiObjEndpointsTSTemplate string

func NewApiObjEndpointsTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__api_obj_endpoints_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetModelFilename":    GetTSModelFilename,
			"GetModelApiFilename": GetTSModelApiFilename,
			"GetObjModelFilename": GetTSModelFilename,
			"GetEnumFilename": func(enumName string) string {
				enum, _ := registry.GetEnum(enumName)
				return GetTSEnumFilename(enum)
			},
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"FormatIndexParams": func(obj types.Object, i types.Index) string {
				var sList []string
				for _, f := range i.Fields {
					varName := utils.LCC(f.Name)
					field, _ := obj.GetField(f.Name)
					fieldType := getTSType(field, registry)
					sList = append(sList, fmt.Sprintf("%s: %s;", varName, fieldType))
				}
				return strings.Join(sList, "\n\t")
			},
			"FormatIndexUrl": func(obj types.Object, i types.Index) string {
				var sList []string
				for _, f := range i.Fields {
					varName := utils.LCC(f.Name)
					sList = append(sList, fmt.Sprintf("%s/${params.%s}", varName, varName))
				}
				return strings.Join(sList, "/")
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getTSDerefType(field, registry)
			},
			"GetFieldDerefRootType": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				return rootType
			},
		})).
		Parse(apiObjEndpointsTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
