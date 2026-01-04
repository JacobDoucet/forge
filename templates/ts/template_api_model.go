package model_template_ts

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed api/model.ts.tmpl
var apiModelTSTemplate string

func NewApiModelTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__obj_api_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{})).
		Parse(apiModelTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
