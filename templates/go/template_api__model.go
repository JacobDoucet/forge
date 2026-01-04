package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed api/model.go.tmpl
var apiModelTemplate string

func NewApiModelGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("http_server_routes").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return "api"
				},
				"GetImports": func(_ types.Object) []string {
					var imports []string
					for _, obj := range registry.ListObjects() {
						if obj.HasCollection() {
							imports = append(imports, registry.GetGoPkgRoot()+templates.GetApiPackageName(obj))
						}
					}
					return imports
				},
				"ListObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasCollection() {
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

	apiTmpl, err := newTmpl().Parse(apiModelTemplate)
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
