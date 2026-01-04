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

//go:embed obj__model/model.go.tmpl
var objModelGoTemplate string

func NewObjModelGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					var hasTimestampField bool
					for _, field := range obj.Fields {
						if rootFieldType, _, _ := field.ResolveRootType(registry); rootFieldType == string(types.FieldTypeTimestamp) {
							hasTimestampField = true
							break
						}
					}

					var imports []string
					if obj.HasCollection() {
						imports = append(imports, packageBsonPrimitive)
						imports = append(imports, packageErrors)
					}
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
				},
				"GetFieldType": func(field types.Field) string {
					return getGoModelType(field, registry)
				},
				"GetHTTPFieldType": func(field types.Field) string {
					return getGoHTTPType(field, registry)
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "Model")
				},
				"AbacToMongoUpdateWhereClause": func(abac types.AbacImpl, obj types.Object, receiverVarName, returnVarName string) string {
					field, _ := obj.GetField(abac.Field)
					if abac.Field == "id" {
						return valueFieldTransform(fieldTransformParams{
							Field:           field,
							Obj:             obj,
							Registry:        registry,
							RefStructName:   "MongoRecord",
							IdTransformType: IdTransformStringToMongo,
							ReturnVarName:   returnVarName,
							ReceiverVarName: receiverVarName,
							FieldName:       receiverVarName + "." + utils.UCC(abac.Name),
							ResolveFieldType: func(field types.Field) string {
								return getGoMongoType(field, registry)
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
					}
					return abacFieldToWhereClause(abac, obj, registry)
				},
				"ModelToMongoRecordTransform": func(field types.Field, obj types.Object, receiverVarName, returnVarName string) string {
					return modelFieldToProjectedFieldTransform(fieldTransformParams{
						Field:             field,
						Obj:               obj,
						Registry:          registry,
						RefStructName:     "MongoRecord",
						IdTransformType:   IdTransformStringToMongo,
						ReturnVarName:     returnVarName,
						ReceiverVarName:   receiverVarName,
						FieldName:         receiverVarName + "." + utils.UCC(field.Name),
						TransformFuncArgs: "projection." + utils.UCC(field.Name) + "Fields",
						ResolveFieldType: func(field types.Field) string {
							return getGoMongoType(field, registry)
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
						IsValueToPointer: true,
					})
				},
				"ModelToHTTPRecordTransform": func(field types.Field, obj types.Object, receiverVarName, returnVarName string) string {
					return modelFieldToProjectedFieldTransform(fieldTransformParams{
						Field:             field,
						Obj:               obj,
						Registry:          registry,
						RefStructName:     "HTTPRecord",
						ReturnVarName:     returnVarName,
						ReceiverVarName:   receiverVarName,
						FieldName:         receiverVarName + "." + utils.UCC(field.Name),
						TransformFuncArgs: "projection." + utils.UCC(field.Name) + "Fields",
						ResolveFieldType: func(field types.Field) string {
							return getGoHTTPType(field, registry)
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
						IsValueToPointer: true,
					})
				},
				"SearchType": func() string {
					return ""
				},
				"GetWhereClause": func(field types.Field) []QueryOption {
					return getWhereClause(field, registry)
				},
				"FormatSearchOptionTypeDef": func(opt QueryOption, obj types.Object) string {
					return opt.FormatModelFieldDef(obj, registry)
				},
				"FormatSearchOptionName": func(opt QueryOption) string {
					return opt.FormatSearchOptionName()
				},
				"FromSearchType": func() string {
					return "WhereClause"
				},
				"ToSearchType": func() string {
					return "MongoWhereClause"
				},
				"ToSelectType": func() string {
					return "Mongo"
				},
				"TransformSelectByField": func(field types.Field, obj types.Object) string {
					return valueFieldTransform(fieldTransformParams{
						Field:           field,
						Obj:             obj,
						Registry:        registry,
						RefStructName:   "MongoWhereClause",
						IdTransformType: IdTransformStringToMongo,
						ReturnVarName:   "to",
						ReceiverVarName: "o",
						FieldName:       "o" + "." + utils.UCC(field.Name),
						ResolveFieldType: func(field types.Field) string {
							return getGoMongoType(field, registry)
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
						RefStructName:   "MongoWhereClause",
						IdTransformType: IdTransformStringToMongo,
						ReturnVarName:   "to",
						ReceiverVarName: "o",
						FieldName:       "o" + "." + opt.FormatSearchOptionName(),
						ResolveFieldType: func(field types.Field) string {
							return getGoMongoType(field, registry)
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
					return ""
				},
				"FormatSortParam": func(field types.IndexField) string {
					return field.FormatSortParam() + " int8"
				},
				"ToSortType": func() string {
					return "Mongo"
				},
				"TransformSortParam": func(field types.IndexField) string {
					return "to." + field.FormatSortParam() + " = s." + field.FormatSortParam()
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	modelTmpl, err := newTmpl().Parse(objModelGoTemplate)
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

	orderTmpl, err := newTmpl().Parse(objSortGoTemplate)
	if err != nil {
		return nil, err
	}

	orderTransformTmpl, err := newTmpl().Parse(objSortTransformGoTemplate)
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
		if err := orderTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err := orderTransformTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func modelFieldToProjectedFieldTransform(params fieldTransformParams) string {
	if params.Field.Name == "id" {
		return strings.Join([]string{
			"if " + params.ReceiverVarName + "." + utils.UCC(params.Field.Name) + " !=\"\" {",
			valueFieldTransform(params),
			"}",
		}, "\n")
	}
	if params.Field.IsRef() {
		return strings.Join([]string{
			"if projection." + utils.UCC(params.Field.Name) + " && " + params.ReceiverVarName + "." + utils.UCC(params.Field.Name) + " !=\"\" {",
			valueFieldTransform(params),
			"}",
		}, "\n")
	}
	return strings.Join([]string{
		"if projection." + utils.UCC(params.Field.Name) + "{",
		valueFieldTransform(params),
		"}",
	}, "\n")
}

func abacFieldToWhereClause(abac types.AbacImpl, obj types.Object, registry types.Registry) string {
	field, _ := obj.GetField(abac.Field)
	mongoType := getGoMongoType(field, registry)

	whereFieldName := utils.UCC(field.Name) + "In"

	assign := "*where." + whereFieldName + " = append(*where." + whereFieldName + ", val)"
	if _, isRef := field.ParseRef(); isRef {
		assign = strings.Join(
			transformStringToObjectId(transformStringToObjectIdParams{
				ElemVar:          "oid",
				FieldName:        utils.LCC(field.Name),
				ReturnVarName:    "where",
				Field:            field,
				ShouldDerefToVar: true,
				TransformElemVarAssignment: func(_ string) string {
					return "append(*where." + whereFieldName + ", oid)"
				},
				TransformReturnField: func(_ string) string {
					return whereFieldName
				},
			}),
			"\n")
	}

	strs := []string{
		"if where." + whereFieldName + "== nil {",
		"where." + whereFieldName + " = &[]" + mongoType + "{}",
		"}",
		utils.LCC(field.Name) + " := role." + utils.UCC(abac.Name),
		"if " + utils.LCC(field.Name) + " != \"\" {",
		assign,
		"}",
	}

	return strings.Join(strs, "\n")
}
