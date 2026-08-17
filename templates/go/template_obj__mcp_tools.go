package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed obj__mcp/tools.go.tmpl
var objMcpToolsTemplate string

const packageMcpSdk = "github.com/modelcontextprotocol/go-sdk/mcp"

func NewObjMCPToolsGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("obj_mcp_tools").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetMCPPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						packageContext,
						packageJson,
						packageMcpSdk,
						registry.GetGoPkgRoot() + "permissions",
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
						registry.GetGoPkgRoot() + templates.GetApiPackageName(obj),
					}
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	toolsTmpl, err := newTmpl().Parse(objMcpToolsTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err := headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := toolsTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, nil
}
