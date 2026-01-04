package model_template_ts

import (
	"bytes"
	_ "embed"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed react/form_state/common.ts.tmpl
var reactFormStateCommonTSTemplate string

func NewReactFormStateCommonTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__react_tanstack_query_obj_api_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{})).
		Parse(reactFormStateCommonTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
