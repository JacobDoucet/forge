package model_template_ts

import (
	"errors"
	"fmt"

	"d3tech.com/platform/templates"
	"d3tech.com/platform/types"
	"d3tech.com/platform/utils"
)

type GenParams struct {
	Registry types.Registry
	OutDir   string
	PkgRoot  string
}

const header = "// This file is auto-generated. DO NOT EDIT.\n\n"

func Gen(params GenParams) ([]templates.OutFile, error) {
	var outFiles []templates.OutFile

	var err error
	genFile := func(relPath string, newGen func(registry types.Registry) (templates.TSGenFunc, error), modelCtx templates.TSTemplateContext) {
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
			Path: params.OutDir + relPath,
			Data: append([]byte(header), file...),
		})
	}

	for _, o := range params.Registry.ListObjects() {
		objCtx := templates.TSTemplateContext{
			Name:   o.Name,
			Object: o,
		}
		genFile("model/"+GetTSModelFilename(o)+".ts", NewModelObjModelTsGenerator, objCtx)
		genFile("model/"+GetTSModelApiFilename(o)+".ts", NewModelObjApiTsGenerator, objCtx)

		genFile("permissions/"+GetTSPermissionsCanAccessFilename(o)+".ts", NewPermissionsObjCanAccessTsGenerator, objCtx)

		if o.HasHTTPMethods() {
			genFile("api/"+GetTSApiEndpointsFilename(o)+".ts", NewApiObjEndpointsTsGenerator, objCtx)
			genFile("react/tanstack-query/"+utils.KC(o.Name)+"-queries.ts", NewReactTanstackQueryObjApiTsGenerator, objCtx)
			genFile("react/form-state/"+GetTSFormStateFilename(o)+".ts", NewReactFormStateObjTsGenerator, objCtx)
			genFile("react/mui/"+utils.KC(o.Name)+"-data-grid.ts", NewReactMUIObjDataGridApiTsGenerator, objCtx)
			genFile("react/mui/"+utils.KC(o.Name)+"-search-selector.tsx", NewReactMUIObjSearchSelectorApiTsGenerator, objCtx)
		}
	}

	for _, e := range params.Registry.ListEnums() {
		genFile("model/"+utils.KC(e.Name)+"-enum.ts", NewModelEnumTsGenerator, templates.TSTemplateContext{
			Enum: e,
		})
	}

	genFile("permissions/actor.ts", NewPermissionsActorTsGenerator, templates.TSTemplateContext{})
	genFile("api/model.ts", NewApiModelTsGenerator, templates.TSTemplateContext{})
	genFile("api/errors.ts", NewApiErrorsTsGenerator, templates.TSTemplateContext{})
	genFile("react/api.tsx", NewReactApiTsGenerator, templates.TSTemplateContext{})
	genFile("react/form-state/common.ts", NewReactFormStateCommonTsGenerator, templates.TSTemplateContext{})
	return outFiles, err
}
