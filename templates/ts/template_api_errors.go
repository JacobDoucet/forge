package model_template_ts

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed api/errors.ts.tmpl
var apiErrorsTSTemplate string

func NewApiErrorsTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__api_errors_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"ListErrors": func() []types.CustomError {
				return registry.ListErrors()
			},
			"ErrorType": func() string {
				var strs []string
				for _, e := range registry.ListErrors() {
					strs = append(strs, "'"+e.Code+"'")
				}
				return strings.Join(strs, "\n  | ")
			},
		})).
		Parse(apiErrorsTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
