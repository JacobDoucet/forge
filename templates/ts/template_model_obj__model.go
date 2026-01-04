package model_template_ts

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed model/obj__model.ts.tmpl
var modelObjModelTSTemplate string

func NewModelObjModelTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__obj_model_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetImportObjs": func(obj types.Object) []types.Object {
				var nestedObjs []types.Object
				objMap := make(map[string]struct{})
				for _, field := range obj.Fields {
					if _, isRef := field.ParseRef(); isRef {
						continue
					}
					rootType, _, _ := field.ResolveRootType(registry)
					nestedObj, ok := registry.Get(rootType)
					if !ok {
						continue
					}
					_, ok = objMap[nestedObj.Name]
					if ok {
						continue
					}
					objMap[nestedObj.Name] = struct{}{}
					nestedObjs = append(nestedObjs, nestedObj)
				}
				sort.Slice(nestedObjs, func(i, j int) bool {
					return nestedObjs[i].Name < nestedObjs[j].Name
				})
				return nestedObjs
			},
			"GetImportEnums": func(obj types.Object) []types.Enum {
				var nestedObjs []types.Enum
				objMap := make(map[string]struct{})
				for _, field := range obj.Fields {
					if _, isRef := field.ParseRef(); isRef {
						continue
					}
					rootType, _, _ := field.ResolveRootType(registry)
					nestedObj, ok := registry.GetEnum(rootType)
					if !ok {
						continue
					}
					_, ok = objMap[nestedObj.Name]
					if ok {
						continue
					}
					objMap[nestedObj.Name] = struct{}{}
					nestedObjs = append(nestedObjs, nestedObj)
				}
				return nestedObjs
			},
			"GetObjModelFilename": GetTSModelFilename,
			"GetEnumFilename":     GetTSEnumFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getTSDerefType(field, registry)
			},
			"ProjectionFieldDef": func(field types.Field) string {
				return getProjectionFieldDef(field, registry)
			},
			"ListReadPermissions": func(obj types.Object) []types.ObjectPermissionsDef {
				return obj.Permissions.Read
			},
			"ListWritePermissions": func(obj types.Object) []types.ObjectPermissionsDef {
				return obj.Permissions.Write
			},
			"FormatIndexFieldSortParam": func(field types.IndexField) string {
				return utils.LCC(field.FormatSortParam())
			},
		})).
		Parse(modelObjModelTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}

func getProjectionFieldDef(field types.Field, registry types.Registry) string {
	if field.Name == "id" {
		return "id?: boolean;"
	}
	rootType, rootClass, _ := field.ResolveRootType(registry)
	if rootClass == types.RootFieldTypePrimitive || rootClass == types.RootFieldTypeEnum {
		return fmt.Sprintf(
			"%s?: boolean;",
			utils.LCC(field.Name),
		)
	}

	_, isRef := field.ParseRef()
	if isRef {
		return fmt.Sprintf("%s?: boolean;", utils.LCC(field.Name))
	}

	refObj, _ := registry.Get(rootType)

	return strings.Join([]string{
		fmt.Sprintf("%s?: boolean;",
			utils.LCC(field.Name),
		),
		fmt.Sprintf("%sFields?: %sProjection;",
			utils.LCC(field.Name),
			utils.UCC(refObj.Name),
		),
	}, "\n\t\t")
}
