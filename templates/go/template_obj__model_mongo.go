package model_template_go

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed obj__model/mongo.go.tmpl
var objModelMongoTemplate string

//go:embed obj__model/search_mongo.tmpl
var objModelSearchMongoTemplate string

func NewObjModelMongoGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("mongo_go").
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
					if hasTimestampField {
						imports = append(imports, packageTime)
					}

					imports = append(imports, packageBson)
					if obj.HasCollection() {
						imports = append(imports, packageBsonPrimitive)
					}
					objFields := obj.ListObjectFields(registry, true)
					for _, ref := range objFields {
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(ref))
					}
					enumField := obj.ListEnumFieldEnums(registry)
					for _, enum := range enumField {
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetEnumPackageName(enum))
					}
					return imports
				},
				"GetBsonFieldTag": func(field types.Field) string {
					return getBsonFieldTag(field)
				},
				"GetFieldType": func(field types.Field) string {
					return getGoMongoType(field, registry)
				},
				"MongoRecordToModelTransform": func(field types.Field, obj types.Object, receiverVarName, returnVarName string) string {
					return transformMongoFieldToModelField(field, obj, registry, receiverVarName, returnVarName)
				},
				"SearchType": func() string {
					return "Mongo"
				},
				"GetWhereClause": func(field types.Field) []QueryOption {
					return getWhereClause(field, registry)
				},
				"FormatSearchOptionTypeDef": func(opt QueryOption, obj types.Object) string {
					return opt.FormatMongoFieldDef(obj, registry)
				},
				"FormatSearchOptionName": func(opt QueryOption) string {
					return opt.FormatSearchOptionName()
				},
				"FormatToMongoQuery": func(opt QueryOption, obj types.Object) string {
					return opt.FormatToMongoQuery(obj)
				},
				"SortType": func() string {
					return "Mongo"
				},
				"FormatSortParam": func(field types.IndexField) string {
					return field.FormatSortParam() + " int8"
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	recordTmpl, err := newTmpl().Parse(objModelMongoTemplate)
	if err != nil {
		return nil, err
	}
	searchTypesTmpl, err := newTmpl().Parse(objSearchTypesTemplate)
	if err != nil {
		return nil, err
	}
	searchMongoTmpl, err := newTmpl().Parse(objModelSearchMongoTemplate)
	if err != nil {
		return nil, err
	}
	sortTmpl, err := newTmpl().Parse(objSortGoTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = recordTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = searchTypesTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = searchMongoTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = sortTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func transformMongoFieldToModelField(field types.Field, obj types.Object, registry types.Registry, receiverVarName, returnVarName string) string {
	return mongoFieldToModelFieldTransform(fieldTransformParams{
		Field:           field,
		Obj:             obj,
		Registry:        registry,
		RefStructName:   "Model",
		IdTransformType: IdTransformMongoToString,
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

func mongoFieldToModelFieldTransform(params fieldTransformParams) string {
	return strings.Join([]string{
		"if " + params.ReceiverVarName + "." + utils.UCC(params.Field.Name) + " != nil {",
		valueFieldTransform(params),
		"}",
	}, "\n")
}

func getBsonFieldTag(field types.Field) string {
	if field.Name == "id" {
		return "_id"
	}
	return field.Name
}

func getProjectionFieldDef(field types.Field, registry types.Registry) string {
	if field.Name == "id" {
		return "Id bool `json:\"id\"`"
	}
	rootType, rootClass, _ := field.ResolveRootType(registry)
	if rootClass == types.RootFieldTypePrimitive || rootClass == types.RootFieldTypeEnum {
		return fmt.Sprintf(
			"%s bool `json:\"%s\"`",
			utils.UCC(field.Name),
			utils.LCC(field.Name),
		)
	}

	_, isRef := field.ParseRef()
	if isRef {
		return fmt.Sprintf("%s bool", utils.UCC(field.Name))
	}

	refObj, _ := registry.Get(rootType)

	return strings.Join([]string{
		fmt.Sprintf("%s bool `json:\"%s\"`",
			utils.UCC(field.Name),
			utils.LCC(field.Name),
		),
		fmt.Sprintf("%sFields %s.Projection `json:\"%sFields,omitempty\"`",
			utils.UCC(field.Name),
			templates.GetModelPackageName(refObj),
			utils.LCC(field.Name),
		),
	}, "\n")
}

func getProjectionDefaultValueDef(field types.Field, registry types.Registry) string {
	if field.Name == "id" {
		return "Id: defaultVal"
	}
	rootType, rootClass, _ := field.ResolveRootType(registry)
	_, isRef := field.ParseRef()
	if rootClass == types.RootFieldTypeObject && !isRef {
		refObj, _ := registry.Get(rootType)
		return strings.Join([]string{
			fmt.Sprintf("%s: defaultVal,", utils.UCC(field.Name)),
			fmt.Sprintf("%sFields: %s.NewProjection(defaultVal)", utils.UCC(field.Name), templates.GetModelPackageName(refObj)),
		}, "\n")
	}

	return fmt.Sprintf("%s: defaultVal", utils.UCC(field.Name))
}

func getProjectionToBSON(field types.Field, registry types.Registry, fieldPath []types.Field) string {
	if len(fieldPath) > 10 {
		var p []string
		for _, f := range fieldPath {
			p = append(p, f.Name)
		}
		panic("projection depth too deep for field " + field.Name + " " + strings.Join(p, "."))
	}

	fieldPath = append(fieldPath, field)
	fieldPathLength := len(fieldPath)
	if field.Name == "id" {
		return "projection[\"_id\"] = 1"
	}

	_, isRef := field.ParseRef()
	rootType, rootClass, _ := field.ResolveRootType(registry)
	if isRef || rootClass == types.RootFieldTypePrimitive || rootClass == types.RootFieldTypeEnum {
		return strings.Join(
			[]string{
				fmt.Sprintf("if p.%s {", fmtFieldPath(fieldPath, func(i int, f types.Field) string {
					suffix := ""
					if i < fieldPathLength-1 {
						suffix = "Fields"
					}
					return utils.UCC(f.Name) + suffix
				})),
				fmt.Sprintf("projection[\"%s\"] = 1", fmtFieldPath(fieldPath, func(_ int, f types.Field) string {
					return getBsonFieldTag(f)
				})),
				"}",
			}, "\n")
	}
	var s []string
	obj, _ := registry.Get(rootType)
	for _, objField := range obj.Fields {
		s = append(
			s,
			[]string{getProjectionToBSON(objField, registry, fieldPath)}...,
		)
	}
	return strings.Join([]string{
		fmt.Sprintf("if p.%s {", fmtFieldPath(fieldPath, func(i int, f types.Field) string {
			suffix := ""
			if i < fieldPathLength-1 {
				suffix = "Fields"
			}
			return utils.UCC(f.Name) + suffix
		})),
		strings.Join(s, "\n"),
		"}",
	}, "\n")
}

func fmtFieldPath(fieldPath []types.Field, transform func(i int, field types.Field) string) string {
	var s []string
	for i, f := range fieldPath {
		s = append(s, transform(i, f))
	}
	return strings.Join(s, ".")
}
