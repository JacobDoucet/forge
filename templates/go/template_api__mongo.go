package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed api/mongo.go.tmpl
var apiMongoTemplate string

func NewApiMongoGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("http_server_routes").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return "api"
				},
				"GetImports": func(_ types.Object) []string {
					imports := []string{
						packageMongo,
					}
					for _, obj := range registry.ListObjects() {
						if obj.HasCollection() {
							imports = append(imports, registry.GetGoPkgRoot()+templates.GetApiPackageName(obj))
						}
					}
					return imports
				},
				"ListMongoObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasCollectionType(types.CollectionTypeMongo) {
							objs = append(objs, obj)
						}
					}
					return objs
				},
				"ListNonMongoObjects": func() []types.Object {
					var objs []types.Object
					for _, obj := range registry.ListObjects() {
						if obj.HasCollection() && !obj.HasCollectionType(types.CollectionTypeMongo) {
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

	apiTmpl, err := newTmpl().Parse(apiMongoTemplate)
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
