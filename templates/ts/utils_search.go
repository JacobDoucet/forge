package model_template_ts

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

func (so *QueryOption) FormatModelFieldDef(obj types.Object, registry types.Registry) string {
	name := so.FormatSearchOptionName()

	asField := types.Field{
		Name: so.Field,
		Type: so.GetSearchOptionType(obj, registry),
	}
	fieldType := getTSType(asField, registry)

	switch so.Operator {
	case SearchOperatorNestedObject:
		rootType, _, _ := asField.ResolveRootType(registry)
		nestedObj, _ := registry.Get(rootType)
		return name + "?: " + utils.UCC(nestedObj.Name) + "SearchQuery;"
	case SearchOperatorIn, SearchOperatorNotIn:
		return name + "?: " + fieldType + "[];"
	case SearchOperatorExists:
		return name + "?: boolean;"
	case SearchOperatorIsEmpty:
		return name + "?: boolean;"
	default:
		return name + "?: " + fieldType + ";"
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

func getQueryOptionMuiDataGridOperators(so QueryOption) []string {
	switch so.Operator {
	case SearchOperatorEqual:
		return []string{"equals", "is"}
	case SearchOperatorNotEqual:
		return []string{"notEquals"}
	case SearchOperatorGreaterThan:
		return []string{"greaterThan", "after"}
	case SearchOperatorGreaterThanOrEqual:
		return []string{"greaterThanOrEqual", "onOrAfter"}
	case SearchOperatorLessThan:
		return []string{"lessThan", "before"}
	case SearchOperatorLessThanOrEqual:
		return []string{"lessThanOrEqual", "onOrBefore"}
	case SearchOperatorIn:
		return []string{"in"}
	case SearchOperatorNotIn:
		return []string{"notIn"}
	case SearchOperatorLike:
		return []string{"contains"}
	case SearchOperatorNotLike:
		return []string{"notContains"}
	case SearchOperatorExists:
		return []string{"isEmpty"}
	default:
		return []string{}
	}
}
