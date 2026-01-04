package model_template_go

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

//go:embed obj__api/with_permissions.go.tmpl
var objModelApiWithPermissionsGoTemplate string

func NewObjModelApiWithPermissionsGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetApiPackageName(obj)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						packageContext,
						registry.GetGoPkgRoot() + "permissions",
						registry.GetGoPkgRoot() + templates.GetModelPackageName(obj),
						registry.GetGoPkgRoot() + "actor_trace",
						registry.GetGoPkgRoot() + "coded_error",
					}

					var addedStrErrPkgs bool
					enumPgsAdded := make(map[string]bool)
					for _, field := range obj.ListFields() {
						if field.Required {
							if !addedStrErrPkgs {
								imports = append(imports, packageStrings)
								imports = append(imports, registry.GetGoPkgRoot()+"coded_error")
							}
							if field.IsEnum(registry) {
								pkgName := registry.GetGoPkgRoot() + "enum_" + utils.SC(field.Type)
								if _, ok := enumPgsAdded[pkgName]; !ok {
									enumPgsAdded[pkgName] = true
									imports = append(imports, pkgName)
								}
							}
						}
					}

					refFields := obj.ListRefFields()
					for _, field := range refFields {
						rootType, _, _ := field.ResolveRootType(registry)
						refObj, _ := registry.Get(rootType)
						imports = append(imports, registry.GetGoPkgRoot()+templates.GetModelPackageName(refObj))
					}
					return imports
				},
				"GetFieldType": func(field types.Field) string {
					return getGoModelType(field, registry)
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "Model")
				},
				"ListActors": func() []types.Object {
					return registry.ListActors()
				},
				"RequiredFieldInvalidCondition": func(field types.Field) string {
					if field.Type == string(types.FieldTypeString) {
						return "obj." + utils.UCC(field.Name) + " == \"\""
					}
					if field.IsEnum(registry) {
						enumPkg := "enum_" + utils.SC(field.Type)
						return "valErr := " + enumPkg + ".Validate(obj." + utils.UCC(field.Name) + "); valErr != nil"
					}
					return ""
				},
				"RequiredFieldInvalidError": func(field types.Field) string {
					if field.Type == string(types.FieldTypeString) {
						return fmt.Sprintf("%s cannot be empty", utils.UCC(field.Name))
					}
					if field.IsEnum(registry) {
						return fmt.Sprintf("%s must be a valid %s", utils.UCC(field.Name), utils.UCC(field.Type))
					}
					return fmt.Sprintf("%s is not a valid %s", utils.UCC(field.Name), field.Type)
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	apiTmpl, err := newTmpl().Parse(objModelApiWithPermissionsGoTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = apiTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return format.Source(buf.Bytes())
	}, err
}
