package model_template_ts

import (
	"bytes"
	_ "embed"
	"sort"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

//go:embed model/obj__api.ts.tmpl
var modelObjApiTSTemplate string

func NewModelObjApiTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__obj_api_ts").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetObjModelFilename":   GetTSModelFilename,
			"GetTSModelApiFilename": GetTSModelApiFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getTSDerefType(field, registry)
			},
			"GetFieldDerefRootType": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				return utils.UCC(rootType)
			},
			"GetEnumFilename": GetTSEnumFilename,
			"GetImportSearchObjs": func(obj types.Object) []types.Object {
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
				sort.Slice(nestedObjs, func(i, j int) bool {
					return nestedObjs[i].Name < nestedObjs[j].Name
				})
				return nestedObjs
			},
			"GetWhereClause": func(field types.Field) []QueryOption {
				return getWhereClause(field, registry)
			},
			"FormatSearchOptionTypeDef": func(opt QueryOption, obj types.Object) string {
				return opt.FormatModelFieldDef(obj, registry)
			},
		})).
		Parse(modelObjApiTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
