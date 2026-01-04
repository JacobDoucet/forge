package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed obj__model/projection.go.tmpl
var objModelProjectionTemplate string

func NewObjModelProjectionGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("mongo_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					var imports []string
					imports = append(imports, packageBson)
					objFields := obj.ListObjectFields(registry, true)
					for _, ref := range objFields {
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(ref))
					}
					return imports
				},
				"ProjectionFieldDef": func(field types.Field) string {
					return getProjectionFieldDef(field, registry)
				},
				"ProjectionDefaultValueDef": func(field types.Field) string {
					return getProjectionDefaultValueDef(field, registry)
				},
				"ProjectionToBSON": func(field types.Field) string {
					return getProjectionToBSON(field, registry, nil)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	recordTmpl, err := newTmpl().Parse(objModelProjectionTemplate)
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
