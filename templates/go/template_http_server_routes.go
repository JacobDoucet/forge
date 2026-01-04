package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed http_server/routes.go.tmpl
var httpServerRoutesTemplate string

func NewHTTPServerRoutesGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("http_server_routes").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return "http_server"
				},
				"GetImports": func(_ types.Object) []string {
					imports := []string{
						packageHttp,
						registry.GetGoPkgRoot() + "api",
						registry.GetGoPkgRoot() + "permissions",
					}
					for _, obj := range registry.ListObjects() {
						if obj.HasHTTPMethods() {
							imports = append(imports, registry.GetGoPkgRoot()+templates.GetHTTPPackageName(obj))
						}
					}
					return imports
				},
				"ListObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasHTTPMethods() {
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

	apiTmpl, err := newTmpl().Parse(httpServerRoutesTemplate)
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
