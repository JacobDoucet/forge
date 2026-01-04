package templates

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

type GoTemplateContext struct {
	Name    string
	Object  types.Object
	Enum    types.Enum
	PkgRoot string
}

type GoGeneratorFunc func(object GoTemplateContext) ([]byte, error)

type TSTemplateContext struct {
	Name   string
	Object types.Object
	Enum   types.Enum
}

type TSGenFunc func(object TSTemplateContext) ([]byte, error)

type KotlinTemplateContext struct {
	Name   string
	Object types.Object
	Enum   types.Enum
}

type KotlinGenFunc func(object KotlinTemplateContext) ([]byte, error)

type TemplateProps struct {
	Name     string
	Registry types.Registry
	PkgRoot  string
	Funcs    template.FuncMap
}

func NewTemplateFuncs(registry types.Registry, customFuncs template.FuncMap) template.FuncMap {
	funcs := template.FuncMap{
		"not": func(v bool) bool {
			return !v
		},
		"and": func(b ...bool) bool {
			if len(b) == 0 {
				return false
			}
			for _, v := range b {
				if !v {
					return false
				}
			}
			return true
		},
		"or": func(b ...bool) bool {
			if len(b) == 0 {
				return false
			}
			for _, v := range b {
				if v {
					return true
				}
			}
			return false
		},
		"dict": func(values ...interface{}) map[string]interface{} {
			m := make(map[string]interface{})
			for i := 0; i < len(values); i += 2 {
				key := values[i].(string)
				m[key] = values[i+1]
			}
			return m
		},
		"concat": func(sep string, values ...string) string {
			return strings.Join(values, sep)
		},
		"SC":  utils.SC,
		"KC":  utils.KC,
		"LCC": utils.LCC,
		"UCC": utils.UCC,
		"COPEN": func() string {
			return "{{"
		},
		"CCLOSE": func() string {
			return "}}"
		},
		"GetModelPackageName": func(obj types.Object) string {
			return GetModelPackageName(obj)
		},
		"GetApiPackageName": func(obj types.Object) string {
			return GetApiPackageName(obj)
		},
		"GetMongoPackageName": func(obj types.Object) string {
			return GetMongoPackageName(obj)
		},
		"GetHTTPPackageName": func(obj types.Object) string {
			return GetHTTPPackageName(obj)
		},
		"ListFields": func(obj types.Object) []types.Field {
			fields := obj.Fields
			if obj.HasCollection() {
				fields = append([]types.Field{
					{
						Name: "id",
						Type: fmt.Sprintf("Ref<%s>", obj.Name),
					},
				}, fields...)
			}
			return fields
		},
		"ListRequiredFields": func(obj types.Object) []types.Field {
			fields := make([]types.Field, 0)
			for _, f := range obj.Fields {
				if f.Required {
					fields = append(fields, f)
				}
			}
			return fields
		},
		"ListFieldsWithoutRefs": func(obj types.Object) []types.Field {
			fields := make([]types.Field, 0)
			for _, f := range obj.Fields {
				if _, isRef := f.ParseRef(); !isRef {
					fields = append(fields, f)
				}
			}
			return fields
		},
		"ListObjectFieldsWithoutRefs": func(obj types.Object) []types.Field {
			fields := make([]types.Field, 0)
			for _, f := range obj.Fields {
				if _, isRef := f.ParseRef(); isRef {
					continue
				}
				rootFieldType, rootFieldClass, _ := f.ResolveRootType(registry)
				if rootFieldClass != types.RootFieldTypeObject {
					continue
				}
				_, ok := registry.Get(rootFieldType)
				if !ok {
					continue
				}
				fields = append(fields, f)
			}
			return fields
		},
		"IsFieldEmbeddedObject": func(field types.Field) bool {
			rootFieldType, rootFieldClass, _ := field.ResolveRootType(registry)
			if rootFieldClass != types.RootFieldTypeObject {
				return false
			}
			_, isRef := field.ParseRef()
			if isRef {
				return false
			}
			_, ok := registry.Get(rootFieldType)
			return ok
		},
		"IsFieldRef": func(field types.Field) bool {
			_, isRef := field.ParseRef()
			return isRef
		},
		"IsFieldEnum": func(field types.Field) bool {
			_, rootFieldClass, _ := field.ResolveRootType(registry)
			return rootFieldClass == types.RootFieldTypeEnum
		},
		"IsFieldList": func(field types.Field) bool {
			_, isList := field.ParseList()
			return isList
		},
		"GetRefFieldObject": func(field types.Field) types.Object {
			rootFieldType, _, _ := field.ResolveRootType(registry)
			obj, _ := registry.Get(rootFieldType)
			return obj
		},
		"GetLookupMethods": func(obj types.Object) []LookupMethod {
			return GetLookupMethods(obj)
		},
		"HasCollection": func(obj types.Object) bool {
			return obj.HasCollection()
		},
		"HasHTTPMethods": func(obj types.Object) bool {
			return obj.HasHTTPMethods()
		},
		"HasAggregation": func(obj types.Object) bool {
			return obj.HasAggregation()
		},
		"ListAggregateFields": func(obj types.Object) []types.AggregateFieldDef {
			return obj.ListAggregateFields()
		},
		"ListGroupByFields": func(obj types.Object) []string {
			return obj.ListGroupByFields()
		},
		// ListGroupByFieldDefs returns the Field definitions for group-by fields
		"ListGroupByFieldDefs": func(obj types.Object) []types.Field {
			var result []types.Field
			for _, fieldName := range obj.ListGroupByFields() {
				if field, ok := obj.GetField(fieldName); ok {
					result = append(result, field)
				}
			}
			return result
		},
		// ListAggregateOnlyFields returns aggregate fields that are NOT in the group-by list
		"ListAggregateOnlyFields": func(obj types.Object) []types.AggregateFieldDef {
			groupBySet := make(map[string]bool)
			for _, g := range obj.ListGroupByFields() {
				groupBySet[g] = true
			}
			var result []types.AggregateFieldDef
			for _, f := range obj.ListAggregateFields() {
				if !groupBySet[f.Field] {
					result = append(result, f)
				}
			}
			return result
		},
		// ListAggregateOnlyFieldDefs returns Field definitions for aggregate fields NOT in group-by
		"ListAggregateOnlyFieldDefs": func(obj types.Object) []types.Field {
			groupBySet := make(map[string]bool)
			for _, g := range obj.ListGroupByFields() {
				groupBySet[g] = true
			}
			var result []types.Field
			for _, f := range obj.ListAggregateFields() {
				if !groupBySet[f.Field] {
					if field, ok := obj.GetField(f.Field); ok {
						result = append(result, field)
					}
				}
			}
			return result
		},
		"ListPrimitiveFields": func(obj types.Object) []types.Field {
			return obj.ListPrimitiveFields(registry)
		},
		"ListEnumFields": func(obj types.Object) []types.Field {
			return obj.ListEnumFields(registry)
		},
		"ListObjectFields": func(obj types.Object) []types.Object {
			return obj.ListObjectFields(registry, false)
		},
		"ListObjectListFields": func(obj types.Object) []types.Object {
			return obj.ListObjectListFields(registry, true)
		},
		"ListRefObjects": func(obj types.Object) []types.Object {
			objs := obj.ListRefObjects()
			sort.Slice(objs, func(i, j int) bool {
				return objs[i].Name < objs[j].Name
			})
			return objs
		},
		"HasObjectFields": func(obj types.Object) bool {
			return len(obj.ListObjectFields(registry, false)) > 0
		},
		"ListRefFields": func(obj types.Object) []types.Field {
			return obj.ListRefFields()
		},
		"ListToManyRefFields": func(obj types.Object) []types.Field {
			return obj.ListToManyRefFields()
		},
		"ListToOneRefFields": func(obj types.Object) []types.Field {
			return obj.ListToOneRefFields()
		},
		"ListParentRefFields": func(obj types.Object) []types.Field {
			return obj.ListParentRefFields()
		},
		"ListChildRefFields": func(obj types.Object) []types.Field {
			return obj.ListChildRefFields()
		},
		"HasRefFields": func(obj types.Object) bool {
			return len(obj.ListRefObjects()) > 0
		},
		"TrimIdSuffix": func(name string) string {
			return TrimIDSuffix(name)
		},
		"GetField": func(name string, obj types.Object) types.Field {
			f, _ := obj.GetField(name)
			return f
		},
		"ListIndexes": func(obj types.Object) []types.Index {
			return ListIndexes(obj)
		},
		"ListIndexFields": func(obj types.Object, index types.Index) []types.Field {
			fields := make([]types.Field, 0)
			for _, f := range index.Fields {
				field, _ := obj.GetField(f.Name)
				fields = append(fields, field)
			}
			return fields
		},
		"ListAllIndexFields": func(obj types.Object) []types.IndexField {
			return ListAllIndexFields(obj)
		},
		"GetRefToModelVarType": func(field types.Field) string {
			return GetRefToModelVarType(field, registry)
		},
		"GetFieldRootType": func(field types.Field) string {
			rootType, _, _ := field.ResolveRootType(registry)
			return rootType
		},
		"GetRefFieldObj": func(field types.Field) types.Object {
			rootType, _, _ := field.ResolveRootType(registry)
			obj, _ := registry.Get(rootType)
			return obj
		},
		"ListRoles": func() []types.PermissionsRbacDef {
			return registry.ListRbacPermissions()
		},
		"GetRbacRoleName": func(permissionDef types.ObjectPermissionsDef) string {
			return "permissions.Role" + utils.UCC(permissionDef.Rbac)
		},
		"ListEnumUniqueIndexFields": func(obj types.Object) []types.Field {
			var enumFields []types.Field
			foundFields := map[string]bool{}
			for _, idx := range obj.Indexes {
				if !idx.Unique {
					continue
				}
				for _, f := range idx.Fields {
					field, _ := obj.GetField(f.Name)
					if _, ok := foundFields[field.Name]; ok {
						continue
					}
					if _, ok := registry.GetEnum(field.Type); ok {
						enumFields = append(enumFields, field)
						foundFields[field.Name] = true
					}
				}
			}
			return enumFields
		},
	}

	for k, v := range customFuncs {
		funcs[k] = v
	}
	return funcs
}

func GetGoEnumPackageDirname(enum types.Enum) string {
	return "enum_" + utils.SC(enum.Name)
}

func GetGoPackageDirname(obj types.Object) string {
	return utils.SC(obj.Name)
}

func GetModelPackageName(obj types.Object) string {
	return utils.SC(obj.Name)
}

func GetEnumPackageName(enum types.Enum) string {
	return "enum_" + utils.SC(enum.Name)
}

func GetApiPackageName(obj types.Object) string {
	return utils.SC(obj.Name) + "_api"
}

func GetMongoPackageName(obj types.Object) string {
	return utils.SC(obj.Name) + "_mongo"
}

func GetHTTPPackageName(obj types.Object) string {
	return utils.SC(obj.Name) + "_http"
}

func GetRefToModelVarType(field types.Field, registry types.Registry) string {
	rootType, _, _ := field.ResolveRootType(registry)
	obj, _ := registry.Get(rootType)
	t := GetModelPackageName(obj) + ".Model"
	if _, isList := field.ParseList(); isList {
		return "[]" + t
	}
	return t
}

func TrimIDSuffix(name string) string {
	if strings.HasSuffix(name, "Id") {
		return name[:len(name)-2]
	}
	if strings.HasSuffix(name, "ID") {
		return name[:len(name)-2]
	}
	if strings.HasSuffix(name, "Ids") {
		return name[:len(name)-3] + "s"
	}
	if strings.HasSuffix(name, "IDs") {
		return name[:len(name)-2] + "s"
	}
	return name
}

type LookupMethod struct {
	Name        string
	IsSelectOne bool
	Index       types.Index
}

func GetLookupMethods(obj types.Object) []LookupMethod {
	var methods []LookupMethod
	for _, idx := range obj.Indexes {
		fieldNames := make([]string, 0)
		fields := make([]types.Field, 0)
		for i, field := range idx.Fields {
			fieldNames = append(fieldNames, utils.UCC(field.Name))
			f, _ := obj.GetField(field.Name)
			fields = append(fields, f)

			isSelectOne := idx.Unique && i == len(idx.Fields)-1

			fieldsConcat := strings.Join(fieldNames, "")
			methodName := fmt.Sprintf("SearchBy%s", fieldsConcat)
			if isSelectOne {
				methodName = fmt.Sprintf("SelectBy%s", fieldsConcat)
			}

			next := LookupMethod{
				Name:        methodName,
				IsSelectOne: isSelectOne,
				Index:       idx,
			}
			methods = append(methods, next)
		}
	}
	return methods
}

func ListIndexes(obj types.Object) []types.Index {
	if !obj.HasCollection() {
		return nil
	}
	return append([]types.Index{{
		Name:   "id",
		Fields: []types.IndexField{{Name: "id"}},
		Unique: true,
	}},
		obj.Indexes...)
}

func ListAllIndexFields(obj types.Object) []types.IndexField {
	fields := make(map[string]types.IndexField)
	for _, idx := range obj.Indexes {
		for _, field := range idx.Fields {
			fields[field.Name] = field
		}
	}
	var result []types.IndexField
	for _, field := range fields {
		result = append(result, field)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}
