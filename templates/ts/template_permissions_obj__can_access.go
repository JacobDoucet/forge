package model_template_ts

import (
	"bytes"
	_ "embed"
	"sort"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed permissions/obj__can_access.ts.tmpl
var permissionsObjCanAccessTSTemplate string

func NewPermissionsObjCanAccessTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("permissions_obj__can_access").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetImportObjs": func(obj types.Object) []types.Object {
				var nestedObjs []types.Object
				objMap := make(map[string]struct{})
				for _, field := range obj.Fields {
					if _, isRef := field.ParseRef(); isRef {
						continue
					}
					rootType, _, _ := field.ResolveRootType(registry)
					nestedObj, ok := registry.Get(rootType)
					if !ok {
						continue
					}
					_, ok = objMap[nestedObj.Name]
					if ok {
						continue
					}
					objMap[nestedObj.Name] = struct{}{}
					nestedObjs = append(nestedObjs, nestedObj)
				}
				//if _, ok := objMap["ActorRole"]; !ok {
				//	actorRole, _ := registry.Get("ActorRole")
				//	nestedObjs = append(nestedObjs, actorRole)
				//}
				sort.Slice(nestedObjs, func(i, j int) bool {
					return nestedObjs[i].Name < nestedObjs[j].Name
				})
				return nestedObjs
			},
			"HasObjectField": func(obj types.Object) bool {
				for _, field := range obj.Fields {
					if _, isRef := field.ParseRef(); isRef {
						continue
					}
					_, rootFieldType, _ := field.ResolveRootType(registry)
					if rootFieldType == types.RootFieldTypeObject {
						return true
					}
				}
				return false
			},
			"IsFieldObject": func(field types.Field) bool {
				_, isRef := field.ParseRef()
				if isRef {
					return false
				}
				_, rootFieldType, _ := field.ResolveRootType(registry)
				return rootFieldType == types.RootFieldTypeObject
			},
			"GetObjModelFilename":       GetTSModelFilename,
			"GetObjPermissionsFilename": GetTSPermissionsCanAccessFilename,
			"GetFieldType": func(field types.Field) string {
				return getTSType(field, registry)
			},
			"ListReadPermissions": func(obj types.Object) []types.ObjectPermissionsDef {
				return obj.Permissions.Read
			},
			"ListWritePermissions": func(obj types.Object) []types.ObjectPermissionsDef {
				return obj.Permissions.Write
			},
		})).
		Parse(permissionsObjCanAccessTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
