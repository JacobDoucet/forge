package model_template_ts

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed api/obj__endpoints.ts.tmpl
var apiObjModelTSTemplate string

func NewApiObjModelTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__api_obj_model_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetObjModelFilename": GetTSModelFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
		})).
		Parse(apiObjModelTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
