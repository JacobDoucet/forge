package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed obj__api/mongo.go.tmpl
var objApiMongoGoTemplate string

func NewObjApiMongoGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
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
						registry.GetGoPkgRoot() + templates.GetMongoPackageName(obj),
						packageErrors,
						packageMongo,
						packageContext,
						packageBsonPrimitive,
					}

					refFields := obj.ListRefFields()
					for _, field := range refFields {
						if _, isList := field.ParseList(); !isList {
							continue
						}
						rootType, _, _ := field.ResolveRootType(registry)
						refObj, _ := registry.Get(rootType)
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
					}
					return imports
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "Model")
				},
				"ToApiModelTransform": func(field types.Field) string {
					return fieldToApiModelTransform(field, registry)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(objApiMongoGoTemplate)
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
