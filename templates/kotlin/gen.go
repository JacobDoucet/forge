package model_template_kotlin

import (
	"errors"
	"fmt"
	"strings"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
)

type GenParams struct {
	Registry types.Registry
	OutDir   string
	PkgRoot  string
}

const header = "// This file is auto-generated. DO NOT EDIT.\n\n"

func Gen(params GenParams) ([]templates.OutFile, error) {
	var outFiles []templates.OutFile

	if !strings.HasPrefix(params.PkgRoot, "com.") || strings.HasSuffix(params.PkgRoot, ".") {
		return nil, fmt.Errorf("invalid PkgRoot: %s", params.PkgRoot)
	}

	outDir := params.OutDir
	pkgSplit := strings.Split(params.PkgRoot, ".")
	for _, part := range pkgSplit {
		outDir += part + "/"
	}

	var err error
	genFile := func(relPath string, newGen func(registry types.Registry) (templates.KotlinGenFunc, error), modelCtx templates.KotlinTemplateContext) {
		genFunc, genErr := newGen(params.Registry)
		if genErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to create generator for %s", relPath), genErr)
			return
		}

		file, genErr := genFunc(modelCtx)
		if genErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to generate %s", relPath), genErr)
			return
		}

		outFiles = append(outFiles, templates.OutFile{
			Path: outDir + relPath,
			Data: append([]byte(header), file...),
		})
	}

	for _, o := range params.Registry.ListObjects() {
		objCtx := templates.KotlinTemplateContext{
			Name:    o.Name,
			Object:  o,
		}
		genFile("model/"+GetKotlinModelFilename(o)+".kt", NewModelObjModelKotlinGenerator, objCtx)
		genFile("model/"+GetKotlinModelApiFilename(o)+".kt", NewModelObjApiKotlinGenerator, objCtx)

		if o.HasHTTPMethods() {
			genFile("api/"+GetKotlinApiEndpointsFilename(o)+".kt", NewApiObjEndpointsKotlinGenerator, objCtx)
		}
	}

	for _, e := range params.Registry.ListEnums() {
		genFile("model/"+GetKotlinEnumFilename(e)+".kt", NewModelEnumKotlinGenerator, templates.KotlinTemplateContext{
			Enum:    e,
		})
	}

	genFile("api/ApiModel.kt", NewApiModelKotlinGenerator, templates.KotlinTemplateContext{})

	return outFiles, err
}
