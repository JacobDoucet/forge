package model_template_kotlin

import (
	"bytes"
	_ "embed"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed api/model.kt.tmpl
var apiModelKotlinTemplate string

func NewApiModelKotlinGenerator(registry types.Registry) (templates.KotlinGenFunc, error) {
	tmpl, err := template.
		New("package__api_model_kotlin").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPkgName": func() string {
				return registry.GetKotlinPkgRoot() + ".api"
			},
		})).
		Parse(apiModelKotlinTemplate)

	return func(ctx templates.KotlinTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyKotlinFile(buf.Bytes()), err
	}, err
}
