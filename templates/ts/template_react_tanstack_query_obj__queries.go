package model_template_ts

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed react/tanstack_query/obj__queries.ts.tmpl
var reactObjApiTSTemplate string

func NewReactTanstackQueryObjApiTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__react_tanstack_query_obj_api_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetObjModelFilename":       GetTSModelFilename,
			"GetObjModelApiFilename":    GetTSModelApiFilename,
			"GetTSApiEndpointsFilename": GetTSApiEndpointsFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"ListSelectFieldsFromParams": func(idx types.Index) string {
				var pList []string
				for _, field := range idx.Fields {
					pList = append(pList, fmt.Sprintf("params.%s", field.Name))
				}
				return strings.Join(pList, ", ")
			},
		})).
		Parse(reactObjApiTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
