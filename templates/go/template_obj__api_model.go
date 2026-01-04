package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"sort"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed obj__api/model.go.tmpl
var objModelApiModelGoTemplate string

func NewObjModelApiModelGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetApiPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						packageContext,
						registry.GetGoPkgRoot() + "permissions",
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
					}

					refFields := obj.ListRefFields()
					for _, field := range refFields {
						rootType, _, _ := field.ResolveRootType(registry)
						refObj, _ := registry.Get(rootType)
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
					}

					// Add time import if aggregation has timestamp group-by fields
					if obj.HasAggregation() {
						for _, fieldName := range obj.ListGroupByFields() {
							if field, ok := obj.GetField(fieldName); ok {
								if field.Type == string(types.FieldTypeTimestamp) {
									imports = append(imports, packageTime)
									break
								}
							}
						}
					}

					sort.Slice(imports, func(i, j int) bool {
						return imports[i] < imports[j]
					})
					return imports
				},
				"GetFieldType": func(field types.Field) string {
					return getGoModelType(field, registry)
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "Model")
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(objModelApiModelGoTemplate)
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
