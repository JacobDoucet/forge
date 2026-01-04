package model_template_go

import (
	"fmt"

	"d3tech.com/platform/templates"
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
	if so.Operator == SearchOperatorExists || so.Operator == SearchOperatorIsEmpty {
		return types.Field{
			Name: so.FormatSearchOptionName(),
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
	fieldType := getGoModelType(asField, registry)

	switch so.Operator {
	case SearchOperatorNestedObject:
		rootType, _, _ := asField.ResolveRootType(registry)
		nestedObj, _ := registry.Get(rootType)
		return name + " *" + templates.GetModelPackageName(nestedObj) + ".WhereClause"
	case SearchOperatorIn, SearchOperatorNotIn:
		return name + " *[]" + fieldType
	case SearchOperatorExists:
		return name + " *bool"
	case SearchOperatorIsEmpty:
		return name + " *bool"
	default:
		return name + " *" + fieldType
	}
}

func (so *QueryOption) FormatHTTPFieldDef(obj types.Object, registry types.Registry) string {
	name := so.FormatSearchOptionName()

	asField := types.Field{
		Name: so.Field,
		Type: so.GetSearchOptionType(obj, registry),
	}
	fieldType := getGoModelType(asField, registry)

	switch so.Operator {
	case SearchOperatorNestedObject:
		rootType, _, _ := asField.ResolveRootType(registry)
		nestedObj, _ := registry.Get(rootType)
		return name + " *" + templates.GetModelPackageName(nestedObj) + ".HTTPWhereClause " + formatJsonTag(name)
	case SearchOperatorIn, SearchOperatorNotIn:
		return name + " *[]" + fieldType + " " + formatJsonTag(name)
	case SearchOperatorExists:
		return name + " *bool " + formatJsonTag(name)
	case SearchOperatorIsEmpty:
		return name + " *bool " + formatJsonTag(name)
	default:
		return name + " *" + fieldType + " " + formatJsonTag(name)
	}
}

func formatJsonTag(fieldName string) string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", utils.LCC(fieldName))
}

func (so *QueryOption) FormatMongoFieldDef(obj types.Object, registry types.Registry) string {
	name := so.FormatSearchOptionName()

	asField := types.Field{
		Name: so.Field,
		Type: so.GetSearchOptionType(obj, registry),
	}
	fieldType := getGoMongoType(asField, registry)

	switch so.Operator {
	case SearchOperatorNestedObject:
		rootType, _, _ := asField.ResolveRootType(registry)
		nestedObj, _ := registry.Get(rootType)
		return name + " *" + templates.GetModelPackageName(nestedObj) + ".MongoWhereClause"
	case SearchOperatorIn, SearchOperatorNotIn:
		return name + " *[]" + fieldType
	case SearchOperatorExists:
		return name + " *bool"
	case SearchOperatorIsEmpty:
		return name + " *bool"
	default:
		return name + " *" + fieldType
	}
}

func (so *QueryOption) FormatSearchOptionName() string {
	return utils.UCC(so.Field) + utils.UCC(string(so.Operator))
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

func (so *QueryOption) FormatToMongoQuery(obj types.Object) string {
	field, _ := obj.GetField(so.Field)
	switch so.Operator {
	case SearchOperatorEqual:
		return "query[\"" + getBsonFieldTag(field) + "\"] = o." + so.FormatSearchOptionName()
	case SearchOperatorNotEqual:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$ne\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorGreaterThan:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$gt\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorGreaterThanOrEqual:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$gte\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorLessThan:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$lt\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorLessThanOrEqual:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$lte\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorIn:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$in\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorNotIn:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$nin\": o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorLike:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$regex\": o." + so.FormatSearchOptionName() + ", \"$options\": \"i\"}"
	case SearchOperatorNotLike:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$not\": bson.M{\"$regex\": o." + so.FormatSearchOptionName() + ", \"$options\": \"i\"}}"
	case SearchOperatorExists:
		return "query[\"" + getBsonFieldTag(field) + "\"] = bson.M{\"$exists\": *o." + so.FormatSearchOptionName() + "}"
	case SearchOperatorNestedObject:
		objLCC := utils.LCC(getBsonFieldTag(field))
		objUCC := utils.UCC(getBsonFieldTag(field))
		return fmt.Sprintf(`%sQuery, err := o.%s.GetQueryParts()
		if err != nil {
			return nil, err
		}
		for _, part := range %sQuery {
			partAsBsonM, ok := part.(bson.M)
			if !ok {
				continue
			}	
			for k, v := range partAsBsonM {
				query["%s."+k] = v
			}
		}`, objLCC, objUCC, objLCC, objLCC)
	case SearchOperatorIsEmpty:
		return `if *o.` + so.FormatSearchOptionName() + ` {
			query["$or"] = bson.A{ 
				bson.M{"` + getBsonFieldTag(field) + `": nil},
				bson.M{"` + getBsonFieldTag(field) + `": bson.A{}},
				bson.M{"` + getBsonFieldTag(field) + `": bson.M{"$exists": false}},
			}
		} else {
			query["` + getBsonFieldTag(field) + `"] = bson.M{
				"$ne": nil,
				"$not": bson.M{"$size": 0},
				"$exists": true,
			}
		}`
	default:
		return ""
	}
}
