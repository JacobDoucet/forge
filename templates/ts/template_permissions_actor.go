package model_template_ts

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed permissions/actor.ts.tmpl
var permissionsActorTSTemplate string

func NewPermissionsActorTsGenerator(registry types.Registry) (templates.TSGenFunc, error) {
	tmpl, err := template.
		New("permissions_obj__can_access").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"ActorType": func() string {
				var t []string
				for _, obj := range registry.ListActors() {
					t = append(t, fmt.Sprintf("'%s'", utils.UCC(obj.Name)))
				}
				return strings.Join(t, " | ")
			},
			"ListCustomPermissions": func() []types.RegisteredCustomPermission {
				return registry.ListCustomPermissions()
			},
			"ListRoleGroups": func() []types.RoleGroupDef {
				return registry.ListRoleGroups()
			},
			"BuildRoleList": func(roleGroup types.RoleGroupDef) string {
				var t []string
				for _, role := range roleGroup.Roles {
					t = append(t, fmt.Sprintf("'%s'", utils.UCC(role)))
				}
				return strings.Join(t, ", ")
			},
		})).
		Parse(permissionsActorTSTemplate)

	return func(ctx templates.TSTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyTSFile(buf.Bytes()), err
	}, err
}
