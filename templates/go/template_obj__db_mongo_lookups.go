package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"sort"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed obj__mongo/model.go.tmpl
var objMongoModelTemplate string

//go:embed obj__mongo/collection.go.tmpl
var objMongoCollectionTemplate string

//go:embed obj__mongo/lookup.go.tmpl
var objMongoLookupTemplate string

//go:embed obj__mongo/search.go.tmpl
var objMongoSearchTemplate string

//go:embed obj__mongo/aggregate.go.tmpl
var objMongoAggregateTemplate string

func NewObjDBMongoCollectionGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	headerTmpl, err := getObjMongoGoTemplate(
		registry,
		func(obj types.Object, registry types.Registry) []string {
			var imports []string
			if len(obj.Indexes) > 0 {
				imports = append(imports, packageContext)
				imports = append(imports, packageMongo)
				imports = append(imports, packageMongoOptions)
				imports = append(imports, packageBson)
			}
			return imports
		},
	).Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	lookupTmpl, err := getObjMongoGoTemplate(
		registry,
		func(obj types.Object, registry types.Registry) []string {
			return []string{}
		},
	).Parse(objMongoCollectionTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = lookupTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func NewObjMongoModelGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	headerTmpl, err := getObjMongoGoTemplate(registry, getObjMongoModelFileImports).Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	lookupTmpl, err := getObjMongoGoTemplate(registry, getObjMongoModelFileImports).Parse(objMongoModelTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = lookupTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func NewObjMongoLookupGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	headerTmpl, err := getObjMongoGoTemplate(registry, getObjMongoLookupFileImports).Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	lookupTmpl, err := getObjMongoGoTemplate(registry, getObjMongoLookupFileImports).Parse(objMongoLookupTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if err = lookupTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func NewObjMongoSearchGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	headerTmpl, err := getObjMongoGoTemplate(registry, getObjMongoSearchFileImports).Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	searchTmpl, err := getObjMongoGoTemplate(registry, getObjMongoSearchFileImports).Parse(objMongoSearchTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if ctx.Object.HasCollection() {
			if err = searchTmpl.Execute(&buf, ctx); err != nil {
				return nil, err
			}
		}
		return format.Source(buf.Bytes())
	}, err
}

func NewObjMongoAggregateGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	headerTmpl, err := getObjMongoGoTemplate(registry, getObjMongoAggregateFileImports).Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}
	aggregateTmpl, err := getObjMongoGoTemplate(registry, getObjMongoAggregateFileImports).Parse(objMongoAggregateTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		if err = headerTmpl.Execute(&buf, ctx); err != nil {
			return nil, err
		}
		if ctx.Object.HasAggregation() {
			if err = aggregateTmpl.Execute(&buf, ctx); err != nil {
				return nil, err
			}
		}
		return format.Source(buf.Bytes())
	}, err
}

func getObjMongoGoTemplate(registry types.Registry, imports func(obj types.Object, registry types.Registry) []string) *template.Template {
	return template.
		New("mongo_go").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPackageName": func(obj types.Object) string {
				return templates.GetMongoPackageName(obj)
			},
			"GetImports": func(obj types.Object) []string {
				return imports(obj, registry)
			},
			"GetFieldType": func(field types.Field) string {
				return getGoMongoType(field, registry)
			},
			"GetFieldRootType": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				return rootType
			},
			"GetFieldCollectionName": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				obj, _ := registry.Get(rootType)
				for _, coll := range obj.Collection {
					if coll.Type == "mongo" {
						return coll.Name
					}
				}
				return "\"NO MONGO COLLECTION FOUND FOR FIELD\""
			},
			"GetFieldDerefType": func(field types.Field) string {
				return getGoDerefType(field, registry, "MongoRecord")
			},
			"GetRefFieldCollection": func(field types.Field) string {
				rootType, _, _ := field.ResolveRootType(registry)
				obj, _ := registry.Get(rootType)
				for _, coll := range obj.Collection {
					if coll.Type == "mongo" {
						return coll.Name
					}
				}
				return "\"NO MONGO COLLECTION FOUND FOR REF\""
			},
			"GetIndexOptions": func(obj types.Object, idx types.Index) string {
				var opts []string
				if idx.Unique && !idx.IncludeNull {
					opts = append(opts, "SetUnique(true)")
					if len(idx.Fields) == 1 {
						if field, ok := obj.GetField(idx.Fields[0].Name); ok {
							opts = append(opts, "SetPartialFilterExpression(bson.D{{Key: \""+field.Name+"\", Value: bson.D{{Key: \"$exists\", Value: true}}}})")
						}
					}

				}
				if idx.Expiration {
					opts = append(opts, "SetExpireAfterSeconds(1)")
				}
				if opts == nil {
					return ""
				}
				return "." + strings.Join(opts, ".\n")
			},
			"GetIndexFieldSortParam": func(field types.IndexField) string {
				return field.FormatSortParam()
			},
			"GetIndexFieldSortMongoKey": func(field types.IndexField) string {
				return field.Name
			},
		}))
}

func getObjMongoLookupFileImports(obj types.Object, registry types.Registry) []string {
	var imports []string

	if obj.HasCollection() {
		imports = append(imports, packageContext)
		imports = append(imports, packageErrors)
		imports = append(imports, packageBson)
		imports = append(imports, packageMongo)
	}

	imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(obj))

	refFields := obj.ListRefFields()
	for _, ref := range refFields {
		rootType, _, _ := ref.ResolveRootType(registry)
		refObj, _ := registry.Get(rootType)
		imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func getObjMongoSelectFileImports(obj types.Object, registry types.Registry) []string {
	var imports []string
	imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(obj))

	if obj.HasCollection() {
		imports = append(imports, packageContext)
		imports = append(imports, packageMongo)
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func getObjMongoSearchFileImports(obj types.Object, registry types.Registry) []string {
	var imports []string

	imports = append(imports, packageContext)
	imports = append(imports, packageMongo)

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func getObjMongoAggregateFileImports(obj types.Object, registry types.Registry) []string {
	var imports []string

	if obj.HasAggregation() {
		imports = append(imports, packageContext)
		imports = append(imports, packageMongo)
		imports = append(imports, packageBson)

		// Check if we need primitive or time imports based on group-by field types
		for _, fieldName := range obj.ListGroupByFields() {
			if field, ok := obj.GetField(fieldName); ok {
				if field.IsRef() {
					imports = append(imports, packageBsonPrimitive)
				}
				if field.Type == string(types.FieldTypeTimestamp) {
					imports = append(imports, packageTime)
				}
			}
		}

		// Add imports for ref field model packages
		refFields := obj.ListRefFields()
		for _, ref := range refFields {
			rootType, _, _ := ref.ResolveRootType(registry)
			refObj, _ := registry.Get(rootType)
			imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
		}
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func getObjMongoModelFileImports(obj types.Object, registry types.Registry) []string {
	var imports []string
	imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(obj))

	refFields := obj.ListRefFields()
	for _, ref := range refFields {
		rootType, _, _ := ref.ResolveRootType(registry)
		refObj, _ := registry.Get(rootType)
		imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
	}

	return imports
}
