package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed coded_error/error.go.tmpl
var codedErrorTemplate string

func NewCodedErrorGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("coded_error").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return "coded_error"
				},
				"GetImports": func(_ types.Object) []string {
					return []string{
						"errors",
						"fmt",
						"strings",
					}
				},
				"ListErrors": func() []types.CustomError {
					return registry.ListErrors()
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(codedErrorTemplate)
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
