package model_template_kotlin

import (
	"fmt"

	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

type SearchOperator string

const (
	SearchOperatorEqual              SearchOperator = "eq"
	SearchOperatorNotEqual           SearchOperator = "ne"
	SearchOperatorGreaterThan        SearchOperator = "gt"
	SearchOperatorGreaterThanOrEqual SearchOperator = "gte"
	SearchOperatorLessThan           SearchOperator = "lt"
	SearchOperatorLessThanOrEqual    SearchOperator = "lte"
	SearchOperatorIn                 SearchOperator = "in"
	SearchOperatorNotIn              SearchOperator = "nin"
	SearchOperatorLike               SearchOperator = "like"
	SearchOperatorNotLike            SearchOperator = "nlike"
	SearchOperatorExists             SearchOperator = "exists"
	SearchOperatorNestedObject       SearchOperator = ""
	SearchOperatorIsEmpty            SearchOperator = "empty"
)

func getWhereClause(field types.Field, registry types.Registry) []QueryOption {
	var options []QueryOption

	if _, isRef := field.ParseRef(); isRef {
		options = append(options, []QueryOption{
			{Field: field.Name, Operator: SearchOperatorEqual},
			{Field: field.Name, Operator: SearchOperatorIn},
			{Field: field.Name, Operator: SearchOperatorNotIn},
			{Field: field.Name, Operator: SearchOperatorExists},
		}...)
		return options
	}

	rootType, rootClass, _ := field.ResolveRootType(registry)

	if rootClass == types.RootFieldTypeObject {
		options = append(options, QueryOption{Field: field.Name, Operator: SearchOperatorNestedObject})
		if _, ok := field.ParseList(); ok {
			options = append(options, QueryOption{Field: field.Name, Operator: SearchOperatorIsEmpty})
		}
		return options
	}

	options = append(options, []QueryOption{
		{Field: field.Name, Operator: SearchOperatorEqual},
		{Field: field.Name, Operator: SearchOperatorNotEqual},
		{Field: field.Name, Operator: SearchOperatorGreaterThan},
		{Field: field.Name, Operator: SearchOperatorGreaterThanOrEqual},
		{Field: field.Name, Operator: SearchOperatorLessThan},
		{Field: field.Name, Operator: SearchOperatorLessThanOrEqual},
		{Field: field.Name, Operator: SearchOperatorIn},
		{Field: field.Name, Operator: SearchOperatorNotIn},
		{Field: field.Name, Operator: SearchOperatorExists},
	}...)

	if rootType == string(types.FieldTypeString) {
		options = append(options, []QueryOption{
			{Field: field.Name, Operator: SearchOperatorLike},
			{Field: field.Name, Operator: SearchOperatorNotLike},
		}...)
	}

	if _, ok := field.ParseList(); ok {
		options = append(options, QueryOption{Field: field.Name, Operator: SearchOperatorIsEmpty})
	}

	return options
}

type QueryOption struct {
	Field    string
	Operator SearchOperator
}

func (so *QueryOption) AsField(obj types.Object, registry types.Registry) types.Field {
	if so.Operator == SearchOperatorNestedObject {
		return types.Field{
			Name: so.Field,
			Type: "bool",
		}
	}
	objField, _ := obj.GetField(so.Field)
	rootType, _, _ := objField.ResolveRootType(registry)
	_, isRef := objField.ParseRef()
	if isRef {
		rootType = "Ref<" + rootType + ">"
	}
	if so.Operator == SearchOperatorIn || so.Operator == SearchOperatorNotIn {
		return types.Field{
			Name: so.FormatSearchOptionName(),
			Type: "List<" + rootType + ">",
		}
	}
	return types.Field{
		Name: so.FormatSearchOptionName(),
		Type: rootType,
	}
}

func (so *QueryOption) FormatKotlinModelFieldDef(obj types.Object, registry types.Registry) string {
	name := so.FormatSearchOptionName()

	asField := types.Field{
		Name: so.Field,
		Type: so.GetSearchOptionType(obj, registry),
	}
	fieldType := getKotlinType(asField, registry)

	switch so.Operator {
	case SearchOperatorNestedObject:
		rootType, _, _ := asField.ResolveRootType(registry)
		nestedObj, _ := registry.Get(rootType)
		return "val " + name + ": " + utils.UCC(nestedObj.Name) + "SearchQuery? = null,"
	case SearchOperatorIn, SearchOperatorNotIn:
		return "val " + name + ": List<" + fieldType + ">? = null,"
	case SearchOperatorExists:
		return "val " + name + ": Boolean? = null,"
	case SearchOperatorIsEmpty:
		return "val " + name + ": Boolean? = null,"
	default:
		return "val " + name + ": " + fieldType + "? = null,"
	}
}

func (so *QueryOption) FormatSearchOptionName() string {
	return utils.LCC(so.Field) + utils.UCC(string(so.Operator))
}

func (so *QueryOption) GetSearchOptionType(object types.Object, registry types.Registry) string {
	field, _ := object.GetField(so.Field)

	_, isRef := field.ParseRef()
	if isRef {
		return fmt.Sprintf("Ref<%s>", field.Type)
	}

	rootType, _, _ := field.ResolveRootType(registry)
	return rootType
}
