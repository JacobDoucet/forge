package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed obj__http/handlers.go.tmpl
var objHttpHandlersTemplate string

func NewObjHTTPHandlersGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetHTTPPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						packageContext,
						packageErrors,
						packageJson,
						packageHttp,
						packageZerolog,
						registry.GetGoPkgRoot() + "permissions",
						registry.GetGoPkgRoot() + "utils",
						registry.GetGoPkgRoot() + "coded_error",
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
						registry.GetGoPkgRoot() + templates.GetApiPackageName(obj),
					}

					if obj.HasAtLeastOneHTTPMethod(types.HttpPUT, types.HttpPATCH, types.HttpPOST) {
						imports = append(imports, packageIo)
					}

					enums := listUniqueIndexEnumFields(registry, obj)
					for _, enum := range enums {
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetEnumPackageName(enum))
					}

					// Add ref model packages for aggregation
					if obj.HasAggregation() {
						refFields := obj.ListRefFields()
						for _, ref := range refFields {
							rootType, _, _ := ref.ResolveRootType(registry)
							refObj, _ := registry.Get(rootType)
							imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
							// Note: We use 'any' type for ref fields in HTTP response to avoid import cycles
						}
					}

					return imports
				},
				"GetModelPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetApiPackageName": func(obj types.Object) string {
					return templates.GetApiPackageName(obj)
				},
				"HasHTTPMethod": func(obj types.Object, method types.HttpMethod) bool {
					return obj.HasHTTPMethod(method)
				},
				"ListQueryStringFields": func(obj types.Object) []types.Field {
					fields := []types.Field{{Name: "id", Type: "string"}}
					for _, field := range obj.Fields {
						if !field.IsPrimitive() {
							continue
						}
						fields = append(fields, field)
					}
					return fields
				},
				"TransformPathValue": func(obj types.Object, fieldName string) string {
					field, _ := obj.GetField(fieldName)
					getPathParam := "r.PathValue(\"" + fieldName + "\")"
					enum, isEnum := registry.GetEnum(field.Type)
					if isEnum {
						return "enum_" + utils.SC(enum.Name) + ".Value(" + getPathParam + ")"
					}
					return getPathParam
				},
				// TypeToGo converts a field type string to Go type
				"TypeToGo": func(fieldType string) string {
					switch types.FieldType(fieldType) {
					case types.FieldTypeBool:
						return "bool"
					case types.FieldTypeString:
						return "string"
					case types.FieldTypeInt:
						return "int"
					case types.FieldTypeInt32:
						return "int32"
					case types.FieldTypeInt64:
						return "int64"
					case types.FieldTypeTimestamp:
						return "string" // Use string for HTTP response (ISO format)
					default:
						// For refs, return string (the ID)
						return "string"
					}
				},
				// GetHTTPPackageName gets the HTTP package name for an object
				"GetHTTPPackageName": func(obj types.Object) string {
					return templates.GetHTTPPackageName(obj)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	httpGetTmpl, err := newTmpl().Parse(objHttpHandlersTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = httpGetTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func listUniqueIndexEnumFields(registry types.Registry, obj types.Object) []types.Enum {
	enums := []types.Enum{}
	foundTypes := map[string]bool{}
	for _, index := range obj.Indexes {
		if !index.Unique {
			continue
		}
		for _, field := range index.Fields {
			f, ok := obj.GetField(field.Name)
			if ok {
				if _, found := foundTypes[f.Type]; found {
					continue
				}
				if enum, isEnum := registry.GetEnum(f.Type); isEnum {
					enums = append(enums, enum)
					foundTypes[f.Type] = true
				}
			}
		}
	}
	return enums
}
