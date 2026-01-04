package model_template_go

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed obj__model/http.go.tmpl
var objModelHttpGoTemplate string

func NewObjModelHTTPGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetHTTPFieldType": func(field types.Field) string {
					return getGoHTTPType(field, registry)
				},
				"GetImports": func(obj types.Object) []string {
					return getHTTPFileImports(obj, registry)
				},
				"HTTPRecordToModelTransform": func(field types.Field, obj types.Object, receiverVarName, returnVarName string) string {
					return transformHTTPFieldToModelField(field, obj, registry, receiverVarName, returnVarName)
				},
				"SelectFieldToModelTransform": func(field types.Field, obj types.Object, receiverVarName, returnVarName string) string {
					return transformHTTPFieldToModelField(field, obj, registry, receiverVarName, returnVarName)
				},
				"IsRootObjectType": func(field types.Field) bool {
					_, isRef := field.ParseRef()
					if isRef {
						return false
					}
					_, rootClass, _ := field.ResolveRootType(registry)
					return rootClass == types.RootFieldTypeObject
				},
				"GetFieldType": func(field types.Field) string {
					return fmt.Sprintf("%s `json:\"%s\"`", getGoHTTPType(field, registry), utils.LCC(field.Name))
				},
				"SearchType": func() string {
					return "HTTP"
				},
				"GetWhereClause": func(field types.Field) []QueryOption {
					return getWhereClause(field, registry)
				},
				"FormatSearchOptionTypeDef": func(opt QueryOption, obj types.Object) string {
					return opt.FormatHTTPFieldDef(obj, registry)
				},
				"FormatSearchOptionName": func(opt QueryOption) string {
					return opt.FormatSearchOptionName()
				},
				"FromSearchType": func() string {
					return "HTTPWhereClause"
				},
				"ToSearchType": func() string {
					return "WhereClause"
				},
				"ToSelectType": func() string {
					return ""
				},
				"SearchOptionAsField": func(opt QueryOption, obj types.Object) types.Field {
					return opt.AsField(obj, registry)
				},
				"TransformSelectByField": func(field types.Field, obj types.Object) string {
					return valueFieldTransform(fieldTransformParams{
						Field:           field,
						Obj:             obj,
						Registry:        registry,
						RefStructName:   "WhereClause",
						ReturnVarName:   "to",
						ReceiverVarName: "o",
						FieldName:       "o" + "." + utils.UCC(field.Name),
						ResolveFieldType: func(field types.Field) string {
							return getGoModelType(field, registry)
						},
						BuildIterVar: func(field types.Field, n int) string {
							return "o" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
						BuildKeyVar: func(field types.Field, n int) string {
							return "oKey" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
						BuildElemVar: func(field types.Field, n int) string {
							return "elem" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
					})
				},
				"TransformSearchField": func(opt QueryOption, obj types.Object) string {
					asField := opt.AsField(obj, registry)
					return valueFieldTransform(fieldTransformParams{
						Field:           asField,
						Obj:             obj,
						Registry:        registry,
						RefStructName:   "WhereClause",
						ReturnVarName:   "to",
						ReceiverVarName: "o",
						FieldName:       "o" + "." + opt.FormatSearchOptionName(),
						ResolveFieldType: func(field types.Field) string {
							return getGoModelType(field, registry)
						},
						BuildIterVar: func(field types.Field, n int) string {
							return "o" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
						BuildKeyVar: func(field types.Field, n int) string {
							return "oKey" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
						BuildElemVar: func(field types.Field, n int) string {
							return "elem" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
						},
						IsPointerToPointer: true,
					})
				},
				"SortType": func() string {
					return "HTTP"
				},
				"FormatSortParam": func(field types.IndexField) string {
					return field.FormatSortParam() + " *int8 `json:\"" + utils.LCC(field.FormatSortParam()) + ",omitempty\"`"
				},
				"ToSortType": func() string {
					return ""
				},
				"TransformSortParam": func(field types.IndexField) string {
					lines := []string{
						"if s." + field.FormatSortParam() + " != nil {",
						"to." + field.FormatSortParam() + " = *s." + field.FormatSortParam(),
						"}",
					}
					return strings.Join(lines, "\n")
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	modelTmpl, err := newTmpl().Parse(objModelHttpGoTemplate)
	if err != nil {
		return nil, err
	}
	searchTypesTmpl, err := newTmpl().Parse(objSearchTypesTemplate)
	if err != nil {
		return nil, err
	}
	searchTransformTmpl, err := newTmpl().Parse(objSearchTypesTransformTemplate)
	if err != nil {
		return nil, err
	}
	sortTmpl, err := newTmpl().Parse(objSortGoTemplate)
	if err != nil {
		return nil, err
	}
	sortTransformTmpl, err := newTmpl().Parse(objSortTransformGoTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = modelTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		if err := searchTypesTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := searchTransformTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := sortTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := sortTransformTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func getHTTPFileImports(obj types.Object, registry types.Registry) []string {
	var hasTimestampField bool
	for _, field := range obj.Fields {
		if rootFieldType, _, _ := field.ResolveRootType(registry); rootFieldType == string(types.FieldTypeTimestamp) {
			hasTimestampField = true
			break
		}
	}

	imports := make([]string, 0)
	if hasTimestampField {
		imports = append(imports, packageTime)
	}

	objFields := obj.ListObjectFields(registry, true)
	for _, field := range objFields {
		imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(field))
	}
	enumField := obj.ListEnumFieldEnums(registry)
	for _, enum := range enumField {
		imports = append(imports, registry.GetGoPkgRoot()+templates.GetEnumPackageName(enum))
	}
	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func transformHTTPFieldToModelField(field types.Field, obj types.Object, registry types.Registry, receiverVarName, returnVarName string) string {
	return httpFieldToModelFieldTransform(fieldTransformParams{
		Field:           field,
		Obj:             obj,
		Registry:        registry,
		RefStructName:   "Model",
		ReturnVarName:   returnVarName,
		ReceiverVarName: receiverVarName,
		FieldName:       receiverVarName + "." + utils.UCC(field.Name),
		ResolveFieldType: func(field types.Field) string {
			return getGoModelType(field, registry)
		},
		BuildIterVar: func(field types.Field, n int) string {
			return receiverVarName + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
		},
		BuildKeyVar: func(field types.Field, n int) string {
			return receiverVarName + "Key" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
		},
		BuildElemVar: func(field types.Field, n int) string {
			return "elem" + fmt.Sprintf("%s%d", utils.LCC(field.Name), n)
		},
		IsPointerToValue: true,
	})
}

func httpFieldToModelFieldTransform(params fieldTransformParams) string {
	return strings.Join([]string{
		"if " + params.ReceiverVarName + "." + utils.UCC(params.Field.Name) + " != nil {",
		valueFieldTransform(params),
		"}",
	}, "\n")
}
