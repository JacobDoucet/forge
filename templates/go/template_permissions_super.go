package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
)

//go:embed permissions/super.go.tmpl
var permissionsSuperTemplate string

func NewPermissionsSuperGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("permission_actor_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					return []string{
						registry.GetGoPkgRoot() + "enum_role",
						registry.GetGoPkgRoot() + "actor_role",
					}
				},
				"ListAbac": func() []types.PermissionsAbacDef {
					return registry.ListAbacPermissions()
				},
				"GetFieldRootType": func(fieldType string) string {
					return getRoleRootFieldType(fieldType, registry)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	actorTmpl, err := newTmpl().Parse(permissionsSuperTemplate)
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
