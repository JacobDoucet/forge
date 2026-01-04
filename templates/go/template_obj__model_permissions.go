package model_template_go

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

//go:embed obj__model/permissions.go.tmpl
var objModelPermissionsTemplate string

func NewObjModelPermissionsGoGenerator(registry types.Registry) (templates.GoGeneratorFunc, error) {
	newTmpl := func() *template.Template {
		return template.
			New("domain_model_go").
			Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
				"GetPackageName": func(obj types.Object) string {
					return templates.GetModelPackageName(obj)
				},
				"GetFieldType": func(field types.Field) string {
					return getGoModelType(field, registry)
				},
				"GetFieldDerefType": func(field types.Field) string {
					return getGoDerefType(field, registry, "Model")
				},
				"GetActorNameStringFmt": func(parts []types.ActorNamePart) string {
					return getActorNameFmt(parts)
				},
				"GetActorNameStringVars": func(parts []types.ActorNamePart) string {
					return getActorNameVars(parts)
				},
				"GetImports": func(obj types.Object) []string {
					imports := []string{
						registry.GetGoPkgRoot() + "permissions",
					}

					appendIfNotExists := func(pkg string) {
						for _, imp := range imports {
							if imp == pkg {
								return
							}
						}
						imports = append(imports, pkg)
					}

					if len(obj.Permissions.Read) > 0 || len(obj.Permissions.Write) > 0 {
						appendIfNotExists(registry.GetGoPkgRoot() + "enum_role")
						appendIfNotExists(registry.GetGoPkgRoot() + "coded_error")
					} else if obj.HasCollection() && len(obj.Abac) > 0 {
						appendIfNotExists(registry.GetGoPkgRoot() + "coded_error")
					}

					if obj.IsActor {
						appendIfNotExists(registry.GetGoPkgRoot() + "actor_role")
					}

					if len(obj.Abac) > 0 {
						appendIfNotExists(packageFmt)
					}

					for _, field := range obj.Fields {
						rootType, rootClass, _ := field.ResolveRootType(registry)
						if rootClass != types.RootFieldTypeObject {
							continue
						}
						_, isRef := field.ParseRef()
						if isRef {
							continue
						}
						embeddedObj, _ := registry.Get(rootType)
						if len(listFieldRBACReadRoles(field)) > 0 || len(listFieldRBACWriteRoles(field)) > 0 {
							appendIfNotExists(registry.GetGoPkgRoot() + templates.GetModelPackageName(embeddedObj))
						}
					}
					return imports
				},
				"HasPermissions": func(obj types.Object) bool {
					return len(obj.Permissions.Read) > 0 || len(obj.Permissions.Write) > 0 || len(listRBACReadRoles(obj)) > 0 || len(listRBACWriteRoles(obj)) > 0
				},
				"HasObjectWritePermissions": func(obj types.Object) bool {
					return len(obj.Permissions.Write) > 0
				},
				"GetObjectPermissionName": func(perm types.ObjectPermissionsDef) string {
					return perm.GetName()
				},
				"ListRBACReadRoles": func(obj types.Object) []string {
					return listRBACReadRoles(obj)
				},
				"HasRbacFieldReadPermissions": func(obj types.Object) bool {
					for _, field := range obj.Fields {
						for _, perm := range field.Permissions.Read {
							if perm.Rbac != "" {
								return true
							}
						}
					}
					return false
				},
				"HasRBACFieldReadRoles": func(field types.Field) bool {
					return len(listFieldRBACReadRoles(field)) > 0
				},
				"ListRBACFieldReadRoles": func(field types.Field) []string {
					return listFieldRBACReadRoles(field)
				},
				"HasRbacFieldWritePermissions": func(obj types.Object) bool {
					for _, field := range obj.Fields {
						for _, perm := range field.Permissions.Write {
							if perm.Rbac != "" {
								return true
							}
						}
					}
					return false
				},
				"HasRBACFieldWriteRoles": func(field types.Field) bool {
					return len(listFieldRBACWriteRoles(field)) > 0
				},
				"ListRBACFieldWriteRoles": func(field types.Field) []string {
					return listFieldRBACWriteRoles(field)
				},
				"FormatAbacListLenCheck": func(perms []types.AbacImpl) string {
					var s []string
					for _, perm := range perms {
						s = append(s, "len("+utils.LCC(perm.Name)+"In) == 0")
					}
					return strings.Join(s, " && ")
				},
				"GetActorIdAbacRuleField": func(obj types.Object) types.Field {
					for _, perm := range obj.Abac {
						if perm.Name == "ActorId" {
							field, _ := obj.GetField(perm.Field)
							return field
						}
					}
					return types.Field{Name: "NOT_DEFINED"}
				},
			}))
	}

	headerTmpl, err := newTmpl().Parse(commonHeaderGoTemplate)
	if err != nil {
		return nil, err
	}

	modelTmpl, err := newTmpl().Parse(objModelPermissionsTemplate)
	if err != nil {
		return nil, err
	}

	return func(ctx templates.GoTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err := headerTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		err = modelTmpl.Execute(&buf, ctx)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
		//return format.Source(buf.Bytes())
	}, err
}

func getActorNameFmt(parts []types.ActorNamePart) string {
	var strParts []string
	for _, part := range parts {
		if part.Field != nil {
			strParts = append(strParts, getFieldStringFmt(*part.Field))
			continue
		}
		strParts = append(strParts, part.Text)
	}
	return strings.Join(strParts, "")
}

func getActorNameVars(parts []types.ActorNamePart) string {
	var fmtParts []string
	for _, part := range parts {
		if part.Field != nil {
			fmtParts = append(fmtParts, "m."+utils.UCC(part.Field.Name))
			continue
		}
	}
	return strings.Join(fmtParts, ", ")
}

func getFieldStringFmt(field types.Field) string {
	switch types.FieldType(field.Type) {
	case types.FieldTypeString:
		return "%s"
	case types.FieldTypeInt, types.FieldTypeInt32, types.FieldTypeInt64:
		return "%d"
	default:
		return "%v"
	}
}

func listRBACReadRoles(obj types.Object) []string {
	var roles []string
	for _, perm := range obj.Permissions.Read {
		if perm.Rbac != "" {
			roles = append(roles, perm.Rbac)
		}
	}
	return roles
}

func listRBACWriteRoles(obj types.Object) []string {
	var roles []string
	for _, perm := range obj.Permissions.Write {
		if perm.Rbac != "" {
			roles = append(roles, perm.Rbac)
		}
	}
	return roles
}

func listFieldRBACReadRoles(field types.Field) []string {
	var roles []string
	for _, perm := range field.Permissions.Read {
		if perm.Rbac != "" {
			roles = append(roles, perm.Rbac)
		}
	}
	return roles
}

func listFieldRBACWriteRoles(field types.Field) []string {
	var roles []string
	for _, perm := range field.Permissions.Write {
		if perm.Rbac != "" {
			roles = append(roles, perm.Rbac)
		}
	}
	return roles
}
