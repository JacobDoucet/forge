package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

//go:embed obj__mongo/delete.go.tmpl
var objMongoDeleteTemplate string

func NewObjMongoDeleteGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
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
					}

					for _, field := range obj.ListChildRefFields() {
						rootType, _, _ := field.ResolveRootType(registry)
						refObj, _ := registry.Get(rootType)
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetMongoPackageName(refObj))
					}

					return imports
				},
				"DeleteChildRef": func(obj types.Object, field types.Field) string {
					rootType, _, _ := field.ResolveRootType(registry)
					refObj, _ := registry.Get(rootType)
					refPkg := templates.GetMongoPackageName(refObj)
					return strings.Join([]string{
						"if err = " + refPkg + ".DeleteBy" + utils.UCC(obj.Name) + "(ctx, db, id); err != nil {",
						"return err",
						"}",
					}, "\n")
				},
				"DeleteManyChildRef": func(obj types.Object, field types.Field) string {
					rootType, _, _ := field.ResolveRootType(registry)
					refObj, _ := registry.Get(rootType)
					refPkg := templates.GetMongoPackageName(refObj)
					return strings.Join([]string{
						"if err = " + refPkg + ".DeleteByMany" + utils.UCC(obj.Name) + "s(ctx, db, ids); err != nil {",
						"return err",
						"}",
					}, "\n")
				},
				"GetFieldRootType": func(field types.Field) string {
					rootType, _, _ := field.ResolveRootType(registry)
					return rootType
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	saveTmpl, err := newTmpl().Parse(objMongoDeleteTemplate)
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
