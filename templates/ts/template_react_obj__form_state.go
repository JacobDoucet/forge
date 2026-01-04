package model_template_ts

import (
	"bytes"
	_ "embed"
	"sort"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed react/form_state/obj.ts.tmpl
var reactFormStateObjTemplate string

func NewReactFormStateObjTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("react_form_state_obj").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetImportObjs": func(obj types.Object) []types.Object {
				var nestedObjs []types.Object
				objMap := make(map[string]struct{})
				for _, field := range obj.Fields {
					if !isFieldFormStateObject(field, registry) {
						continue
					}
					rootType, _, _ := field.ResolveRootType(registry)
					nestedObj, _ := registry.Get(rootType)
					_, ok := objMap[nestedObj.Name]
					if ok {
						continue
					}
					objMap[nestedObj.Name] = struct{}{}
					nestedObjs = append(nestedObjs, nestedObj)
				}
				sort.Slice(nestedObjs, func(i, j int) bool {
					return nestedObjs[i].Name < nestedObjs[j].Name
				})
				return nestedObjs
			},
			"HasObjectField": func(obj types.Object) bool {
				for _, field := range obj.Fields {
					if isFieldFormStateObject(field, registry) {
						return true
					}
				}
				return false
			},
			"IsFieldObject": func(field types.Field) bool {
				return isFieldFormStateObject(field, registry)
			},
			"GetObjModelFilename":       GetTSModelFilename,
			"GetObjPermissionsFilename": GetTSPermissionsCanAccessFilename,
			"GetObjFormStateFilename":   GetTSFormStateFilename,
			"GetTSApiEndpointsFilename": GetTSApiEndpointsFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
		})).
		Parse(reactFormStateObjTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}

func isFieldFormStateObject(field types.Field, registry types.Registry) bool {
	if _, isRef := field.ParseRef(); isRef {
		return false
	}
	if _, isList := field.ParseList(); isList {
		return false
	}
	if _, isKeyVal := field.ParseKeyVal(); isKeyVal {
		return false
	}
	_, rootFieldType, _ := field.ResolveRootType(registry)
	return rootFieldType == types.RootFieldTypeObject
}
