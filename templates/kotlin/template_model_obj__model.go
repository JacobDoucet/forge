package model_template_kotlin

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

//go:embed model/obj__model.kt.tmpl
var modelObjModelKotlinTemplate string

func NewModelObjModelKotlinGenerator(registry types.Registry) (templates.KotlinGenFunc, error) {
	tmpl, err := template.
		New("package__obj_model_kotlin").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPkgName": func() string {
				return registry.GetKotlinPkgRoot() + ".model"
			},
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
			"GetObjModelFilename": GetKotlinModelFilename,
			"GetEnumFilename":     GetKotlinEnumFilename,
			"GetFieldType": func(field types.Field) string {
				return getKotlinType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getKotlinDerefType(field, registry)
			},
			"ProjectionFieldDef": func(field types.Field) string {
				return getKotlinProjectionFieldDef(field, registry)
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
		Parse(modelObjModelKotlinTemplate)

	return func(ctx templates.KotlinTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyKotlinFile(buf.Bytes()), err
	}, err
}

func getKotlinProjectionFieldDef(field types.Field, registry types.Registry) string {
	if field.Name == "id" {
		return "val id: Boolean? = null,"
	}
	rootType, rootClass, _ := field.ResolveRootType(registry)
	if rootClass == types.RootFieldTypePrimitive || rootClass == types.RootFieldTypeEnum {
		return fmt.Sprintf(
			"val %s: Boolean? = null,",
			utils.LCC(field.Name),
		)
	}

	_, isRef := field.ParseRef()
	if isRef {
		return fmt.Sprintf("val %s: Boolean? = null,", utils.LCC(field.Name))
	}

	refObj, _ := registry.Get(rootType)

	return strings.Join([]string{
		fmt.Sprintf("val %s: Boolean? = null,",
			utils.LCC(field.Name),
		),
		fmt.Sprintf("val %sFields: %sProjection? = null,",
			utils.LCC(field.Name),
			utils.UCC(refObj.Name),
		),
	}, "\n    ")
}
