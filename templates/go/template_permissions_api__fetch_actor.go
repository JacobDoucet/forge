package model_template_go

import (
	"bytes"
	_ "embed"
	"go/format"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

//go:embed permissions_api/fetch_actor.go.tmpl
var permissionsApiFetchActorTemplate string

func NewPermissionsApiFetchActorGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("permission_api_fetch_actor_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return "permissions_api"
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						packageContext,
						packageFmt,
						registry.GetGoPkgRoot() + "permissions",
						registry.GetGoPkgRoot() + "coded_error",
					}

					for _, actorObj := range registry.ListActors() {
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(actorObj))
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetApiPackageName(actorObj))
					}

					return imports
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

	actorTmpl, err := newTmpl().Parse(permissionsApiFetchActorTemplate)
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
