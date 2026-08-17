package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed mcp/register.go.tmpl
var mcpRegisterTemplate string

func NewMCPRegisterGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("mcp_register").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(_ types.Object) string {
					return "mcp_register"
				},
				"GetImports": func(_ types.Object) []string {
					imports := []string{
						packageContext,
						packageMcpSdk,
						registry.GetGoPkgRoot() + "api",
						registry.GetGoPkgRoot() + "permissions",
					}
					for _, obj := range registry.ListObjects() {
						if obj.HasMCPMethods() {
							imports = append(imports, registry.GetGoPkgRoot()+templates.GetMCPPackageName(obj))
						}
					}
					return imports
				},
				"ListObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasMCPMethods() {
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

	regTmpl, err := newTmpl().Parse(mcpRegisterTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err := headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := regTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, nil
}
