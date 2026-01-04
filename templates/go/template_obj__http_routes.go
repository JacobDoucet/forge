package model_template_go

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed obj__http/routes.go.tmpl
var objHttpRoutesTemplate string

func NewObjHTTPRoutesGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetHTTPPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						packageHttp,
						packageStrings,
					}
				},
				"HasHTTPMethod": func(obj types.Object, method types.HttpMethod) bool {
					return obj.HasHTTPMethod(method)
				},
				"FormatParam": func(i types.Index) string {
					var sList []string
					for _, f := range i.Fields {
						s := utils.LCC(f.Name)
						sList = append(sList, fmt.Sprintf("%s/{%s}", s, s))
					}
					return strings.Join(sList, "/")
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	httpGetTmpl, err := newTmpl().Parse(objHttpRoutesTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = httpGetTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}
