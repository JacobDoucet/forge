package model_template_go

import (
	"errors"
	"fmt"

	"github.com/JacobDoucet/forge/templates"
	"github.com/JacobDoucet/forge/types"
	"github.com/JacobDoucet/forge/utils"
)

type GenParams struct {
	Registry types.Registry
	OutDir   string
	PkgRoot  string
}

func Gen(params GenParams) ([]templates.OutFile, error) {
	var outFiles []templates.OutFile

	var err error
	genFile := func(relPath string, newGen func(registry types.Registry) (templates.GoGeneratorFunc, error), modelCtx templates.GoTemplateContext) {
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
			Data: file,
		})
	}

	for _, e := range params.Registry.ListEnums() {
		genFile(templates.GetGoEnumPackageDirname(e)+"/value.go", NewEnumNameValueGoGenerator, templates.GoTemplateContext{
			Name:   e.Name,
			Enum:   e,
			Object: types.Object{Name: "enum" + e.Name},
		})
	}

	for _, o := range params.Registry.ListObjects() {
		objCtx := templates.GoTemplateContext{
			Name:    o.Name,
			Object:  o,
			PkgRoot: params.PkgRoot,
		}
		genFile(templates.GetGoPackageDirname(o)+"/model.go", NewObjModelGoGenerator, objCtx)
		genFile(templates.GetGoPackageDirname(o)+"/projection.go", NewObjModelProjectionGoGenerator, objCtx)
		genFile(templates.GetGoPackageDirname(o)+"/mongo.go", NewObjModelMongoGoGenerator, objCtx)
		genFile(templates.GetGoPackageDirname(o)+"/http.go", NewObjModelHTTPGoGenerator, objCtx)

		if o.HasCollection() {
			genFile(templates.GetGoPackageDirname(o)+"/permissions.go", NewObjModelPermissionsGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_api/with_permissions.go", NewObjModelApiWithPermissionsGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_api/model.go", NewObjModelApiModelGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_api/unimplemented.go", NewObjModelApiUnimplementedGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_api/http.go", NewObjModelApiHTTPGoGenerator, objCtx)
			if utils.LCC(o.Name) != "event" {
				genFile(templates.GetGoPackageDirname(o)+"/events.go", NewObjModelEventsGoGenerator, objCtx)
			}
		}
		if o.HasCollectionType(types.CollectionTypeMongo) {
			genFile(templates.GetGoPackageDirname(o)+"_mongo/collection.go", NewObjDBMongoCollectionGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_mongo/model.go", NewObjMongoModelGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_mongo/lookup.go", NewObjMongoLookupGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_mongo/search.go", NewObjMongoSearchGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_mongo/save.go", NewObjMongoSaveGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_mongo/delete.go", NewObjMongoDeleteGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_api/mongo.go", NewObjApiMongoGoGenerator, objCtx)
			if o.HasAggregation() {
				genFile(templates.GetGoPackageDirname(o)+"_mongo/aggregate.go", NewObjMongoAggregateGoGenerator, objCtx)
			}
		}

		if o.HasHTTPMethods() {
			genFile(templates.GetGoPackageDirname(o)+"_http/handlers.go", NewObjHTTPHandlersGoGenerator, objCtx)
			genFile(templates.GetGoPackageDirname(o)+"_http/routes.go", NewObjHTTPRoutesGoGenerator, objCtx)
		}
	}

	genFile("permissions/actor.go", NewPermissionsActorGenerator, templates.GoTemplateContext{
		Object: types.Object{Name: "permissions"},
	})
	genFile("permissions/super.go", NewPermissionsSuperGenerator, templates.GoTemplateContext{
		Object: types.Object{Name: "permissions"},
	})

	genFile("permissions_api/fetch_actor.go", NewPermissionsApiFetchActorGenerator, templates.GoTemplateContext{})

	genFile("http_server/routes.go", NewHTTPServerRoutesGoGenerator, templates.GoTemplateContext{})
	genFile("api/model.go", NewApiModelGoGenerator, templates.GoTemplateContext{})
	genFile("api/mongo.go", NewApiMongoGoGenerator, templates.GoTemplateContext{})

	genFile("coded_error/error.go", NewCodedErrorGoGenerator, templates.GoTemplateContext{})
	genFile("utils/conv.go", NewUtilsConvGoGenerator, templates.GoTemplateContext{})

	genFile("event/interface.go", NewEventInterfaceGoGenerator, templates.GoTemplateContext{})

	return outFiles, err
}
