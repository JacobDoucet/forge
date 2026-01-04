package model_template_kotlin

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

//go:embed api/obj__endpoints.kt.tmpl
var apiObjEndpointsKotlinTemplate string

func NewApiObjEndpointsKotlinGenerator(registry types.Registry) (templates.KotlinGenFunc, error) {
	tmpl, err := template.
		New("package__api_obj_endpoints_kotlin").
		Funcs(templates.NewTemplateFuncs(registry, template.FuncMap{
			"GetPkgName": func() string {
				return registry.GetKotlinPkgRoot() + ".api"
			},
			"GetModelPkgNamePrefix": func() string {
				return registry.GetKotlinPkgRoot() + ".model."
			},
			"GetModelFilename":    GetKotlinModelFilename,
			"GetModelApiFilename": GetKotlinModelApiFilename,
			"FormatKotlinIndexParams": func(obj types.Object, i types.Index) string {
				var sList []string
				for _, f := range i.Fields {
					varName := utils.LCC(f.Name)
					field, _ := obj.GetField(f.Name)
					fieldType := getKotlinType(field, registry)
					sList = append(sList, fmt.Sprintf("val %s: %s,", varName, fieldType))
				}
				return strings.Join(sList, "\n\t\t")
			},
			"FormatKotlinIndexUrl": func(obj types.Object, i types.Index) string {
				var sList []string
				for _, f := range i.Fields {
					varName := utils.LCC(f.Name)
					sList = append(sList, fmt.Sprintf("%s/${params.%s}", varName, varName))
				}
				return strings.Join(sList, "/")
			},
		})).
		Parse(apiObjEndpointsKotlinTemplate)

	return func(ctx templates.KotlinTemplateContext) ([]byte, error) {
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, ctx)
		return tidyKotlinFile(buf.Bytes()), err
	}, err
}

func GetKotlinApiEndpointsFilename(obj types.Object) string {
	return utils.UCC(obj.Name) + "ApiClient"
}
