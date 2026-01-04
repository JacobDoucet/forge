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

//go:embed permissions/actor.go.tmpl
var permissionsActorTemplate string

func NewPermissionsActorGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("permission_actor_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						packageErrors,
						packageTime,
						registry.GetGoPkgRoot() + "enum_role",
						registry.GetGoPkgRoot() + "actor_role",
						registry.GetGoPkgRoot() + "actor_trace",
					}
				},
				"ListAbac": func() []types.PermissionsAbacDef {
					return registry.ListAbacPermissions()
				},
				"GetFieldRootType": func(fieldType string) string {
					return getRoleRootFieldType(fieldType, registry)
				},
				"GetRoleArgs": func(role types.PermissionsRbacDef) string {
					var args []string
					for _, abacName := range role.Abac {
						abac, _ := registry.GetAbacPermission(abacName)
						fieldType := getRoleRootFieldType(abac.FieldType, registry)
						args = append(args, fmt.Sprintf("%s %s", utils.LCC(abacName), fieldType))
					}
					return strings.Join(args, ", ")
				},
				"ListCustomPermissions": func() []types.RegisteredCustomPermission {
					return registry.ListCustomPermissions()
				},
				"ListActorObjects": func() []types.Object {
					return registry.ListActors()
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	actorTmpl, err := newTmpl().Parse(permissionsActorTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = actorTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}

func getRoleRootFieldType(fieldType string, registry types.Registry) string {
	rootType, fieldClass, _ := types.Field{
		Type: fieldType,
	}.ResolveRootType(registry)
	if fieldClass == types.RootFieldTypeObject {
		return "string"
	}
	return rootType
}
