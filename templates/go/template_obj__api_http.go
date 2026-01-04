package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed obj__api/http.go.tmpl
var objModelApiHTTPGoTemplate string

func NewObjModelApiHTTPGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetApiPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
						packageErrors,
					}

					refFields := obj.ListRefFields()
					for _, field := range refFields {
						rootType, _, _ := field.ResolveRootType(registry)
						refObj, _ := registry.Get(rootType)
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
					}
					return imports
				},
				"GetFieldType": func(field types.Field) string {
					return getGoHTTPType(field, registry)
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "HTTPRecord")
				},
				"ToApiModelTransform": func(field types.Field) string {
					return fieldToApiModelTransform(field, registry)
				},
				"ToHTTPRecordTransform": func(field types.Field) string {
					return fieldToHTTPApiModelTransform(field, registry)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(objModelApiHTTPGoTemplate)
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
