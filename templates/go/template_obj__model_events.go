package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed obj__model/events.go.tmpl
var objModelEventsTemplate string

func NewObjModelEventsGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("mongo_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						packageErrors,
						registry.GetGoPkgRoot() + "enum_model",
						registry.GetGoPkgRoot() + "event_subject",
					}
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	recordTmpl, err := newTmpl().Parse(objModelEventsTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = recordTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}
