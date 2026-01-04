package model_template_ts

import (
	"bytes"
	_ "embed"
	"sort"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed react/mui/obj__search_selector.tsx.tmpl
var reactMUIObjSearchSelectorTSTemplate string

func NewReactMUIObjSearchSelectorApiTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("package__react_obj__search_selector").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetObjModelFilename": GetTSModelFilename,
			"GetObjApiFilename":   GetTSModelApiFilename,
			"GetWhereClause": func(field types.Field) []QueryOption {
				return getWhereClause(field, registry)
			},
			"FormatSearchOptionFieldName": func(opt QueryOption) string {
				return opt.FormatSearchOptionName()
			},
			"GetOperators": func(opt QueryOption) []string {
				return getQueryOptionMuiDataGridOperators(opt)
			},
			"FormatSearchOptionTransform": func(opt QueryOption, obj types.Object) string {
				name := opt.FormatSearchOptionName()

				asField := types.Field{
					Name: opt.Field,
					Type: opt.GetSearchOptionType(obj, registry),
				}
				fieldType := getTSType(asField, registry)

				switch opt.Operator {
				case SearchOperatorNestedObject:
					// rootType, _, _ := asField.ResolveRootType(registry)
					// nestedObj, _ := registry.Get(rootType)
					// TODO
					return "null"
				case SearchOperatorIn, SearchOperatorNotIn:
					return name + "?: " + fieldType + "[];"
				case SearchOperatorExists:
					return name + "?: boolean;"
				case SearchOperatorIsEmpty:
					return name + "?: boolean;"
				default:
					return name + "?: " + fieldType + ";"
				}
			},
			"UndefinedIfValueIsUnset": func(opt QueryOption, obj types.Object, valueExpression string) string {
				asField := types.Field{
					Name: opt.Field,
					Type: opt.GetSearchOptionType(obj, registry),
				}
				fieldType := getTSType(asField, registry)

				switch opt.Operator {
				case SearchOperatorNestedObject:
					// rootType, _, _ := asField.ResolveRootType(registry)
					// nestedObj, _ := registry.Get(rootType)
					// TODO
					return "null"
				case SearchOperatorIn, SearchOperatorNotIn:
					return "!" + valueExpression + "?.length ? undefined : "
				case SearchOperatorExists:
					return valueExpression + " == null ? undefined : "
				case SearchOperatorIsEmpty:
					return valueExpression + " == null ? undefined : " + valueExpression
				default:
					switch fieldType {
					case "string":
						return "!" + valueExpression + " ? undefined : "
					case "number", "boolean", "Date":
						return valueExpression + " == null ? undefined : "
					default:
						return ""
					}
				}
			},
			"ListRefFieldObjects": func(obj types.Object) []types.Object {
				m := make(map[string]struct{})
				for _, field := range obj.Fields {
					if name, ok := field.ParseRef(); ok {
						m[name] = struct{}{}
					}
				}
				objs := make([]types.Object, 0, len(m))
				for name := range m {
					if refObj, ok := registry.Get(name); ok {
						objs = append(objs, refObj)
					}
				}

				sort.Slice(objs, func(i, j int) bool {
					return objs[i].Name < objs[j].Name
				})
				return objs
			},
			"GetColumnType": func(field types.Field) string {
				fieldRootType, rootType, _ := field.ResolveRootType(registry)
				if rootType == types.RootFieldTypePrimitive {
					tsType := strings.TrimSuffix(getTSType(field, registry), "[]")
					if fieldRootType == string(types.FieldTypeTimestamp) {
						return `"date"`
					}
					return `"` + tsType + `"`
				}
				return "undefined"
			},
			"GetColumnTSType": func(field types.Field) string {
				fieldRootType, rootType, _ := field.ResolveRootType(registry)
				if rootType == types.RootFieldTypePrimitive {
					tsType := strings.TrimSuffix(getTSType(field, registry), "[]")
					if fieldRootType == string(types.FieldTypeTimestamp) {
						return "Date"
					}
					return tsType
				}
				return "string"
			},
			"GetFieldEnum": func(field types.Field) types.Enum {
				if rootType, rootClass, _ := field.ResolveRootType(registry); rootClass == types.RootFieldTypeEnum {
					enum, _ := registry.GetEnum(rootType)
					return enum
				}
				return types.Enum{}
			},
		})).
		Parse(reactMUIObjSearchSelectorTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
