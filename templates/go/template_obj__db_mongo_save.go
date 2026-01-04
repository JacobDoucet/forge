package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed obj__mongo/save.go.tmpl
var objMongoSaveTemplate string

func NewObjMongoSaveGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetMongoPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						packageContext,
						packageMongo,
						packageBson,
						packageBsonPrimitive,
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
					}

					for _, abac := range obj.Abac {
						if abac.Field == "id" {
							continue
						}
						field, _ := obj.GetField(abac.Field)
						_, isRef := field.ParseRef()
						if isRef {
							imports = append(imports, packageBson)
							break
						}
					}

					return imports
				},
				"GetFieldType": func(field types.Field) string {
					return getGoMongoType(field, registry)
				},
				"GetBsonFieldTag": func(field types.Field) string {
					return getBsonFieldTag(field)
				},
				"ListImmutableFields": func(obj types.Object) []types.Field {
					return obj.ListImmutableFields(registry)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	saveTmpl, err := newTmpl().Parse(objMongoSaveTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = saveTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}
