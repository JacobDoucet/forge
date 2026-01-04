package model_template_kotlin

import (
	"bytes"
	_ "embed"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed model/obj__api.kt.tmpl
var modelObjApiKotlinTemplate string

func NewModelObjApiKotlinGenerator(registry types.Registry) (templates.KotlinGenFunc, error) {
	tmpl, err := template.
		New("package__obj_api_kotlin").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPkgName": func() string {
				return registry.GetKotlinPkgRoot() + ".model"
			},
			"GetObjModelFilename":       GetKotlinModelFilename,
			"GetKotlinModelApiFilename": GetKotlinModelApiFilename,
			"GetEnumFilename":           GetKotlinEnumFilename,
			"GetImportEnums": func(obj types.Object) []types.Enum {
				var nestedEnums []types.Enum
				enumMap := make(map[string]struct{})
				for _, field := range obj.Fields {
					if _, isRef := field.ParseRef(); isRef {
						continue
					}
					rootType, _, _ := field.ResolveRootType(registry)
					nestedEnum, ok := registry.GetEnum(rootType)
					if !ok {
						continue
					}
					_, ok = enumMap[nestedEnum.Name]
					if ok {
						continue
					}
					enumMap[nestedEnum.Name] = struct{}{}
					nestedEnums = append(nestedEnums, nestedEnum)
				}
				return nestedEnums
			},
			"GetImportSearchObjs": func(obj types.Object) []types.Object {
				// For now, return empty slice - implement search objects later
				return []types.Object{}
			},
			"GetWhereClause": func(field types.Field) []QueryOption {
				return getWhereClause(field, registry)
			},
			"GetFieldType": func(field types.Field) string {
				return getKotlinType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getKotlinDerefType(field, registry)
			},
			"GetFieldDerefRootType": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				return rootType
			},
			"ProjectionFieldDef": func(field types.Field) string {
				return getKotlinProjectionFieldDef(field, registry)
			},
			"FormatKotlinSearchOptionTypeDef": func(opt QueryOption, obj types.Object) string {
				return opt.FormatKotlinModelFieldDef(obj, registry)
			},
		})).
		Parse(modelObjApiKotlinTemplate)

	return func(ctx templates.KotlinTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyKotlinFile(buf.Bytes()), err
	}, err
}

func GetKotlinModelApiFilename(obj types.Object) string {
	return GetKotlinModelFilename(obj) + "Api"
}
