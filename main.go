package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"d3tech.com/platform/templates"
	model_template_go "d3tech.com/platform/templates/go"
	model_template_kotlin "d3tech.com/platform/templates/kotlin"
	model_template_ts "d3tech.com/platform/templates/ts"
	"d3tech.com/platform/types"
	"gopkg.in/yaml.v3"
)

type ModelFile struct {
	Objects          []types.Object          `yaml:"objects"`
	Enums            []types.Enum            `yaml:"enums"`
	Permissions      types.PermissionsDef    `yaml:"permissions"`
	Errors           []types.CustomError     `yaml:"errors"`
	Events           []string                `yaml:"events"`
	EventPermissions types.ObjectPermissions `yaml:"eventPermissions"`
}

func main() {
	var goOutDir, goPkgRoot, tsOutDir, kotlinOutDir, kotlinPkgRoot, specDir string
	flag.StringVar(&goOutDir, "goOutDir", "", "out directory of generated files")
	flag.StringVar(&goPkgRoot, "goPkgRoot", "", "root package of generated files")
	flag.StringVar(&tsOutDir, "tsOutDir", "", "out directory of generated ts files")
	flag.StringVar(&kotlinOutDir, "kotlinOutDir", "", "out directory of generated kotlin files")
	flag.StringVar(&kotlinPkgRoot, "kotlinPkgRoot", "", "root package of generated kotlin files")
	flag.StringVar(&specDir, "specDir", "", "directory containing YAML spec files")

	flag.Parse()

	shouldGenGo := goOutDir != "" || goPkgRoot != ""
	shouldGenTS := tsOutDir != ""
	shouldGenKotlin := kotlinOutDir != "" || kotlinPkgRoot != ""

	goOutDir = sanitizePath(goOutDir)
	goPkgRoot = sanitizePath(goPkgRoot)
	tsOutDir = sanitizePath(tsOutDir)
	kotlinOutDir = sanitizePath(kotlinOutDir)

	// ensure out dir
	err := os.MkdirAll(goOutDir, 0755)
	if err != nil {
		panic(err)
	}

	// Check if spec directory exists
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		panic(fmt.Errorf("spec directory %s does not exist", specDir))
	}

	// Read files from the spec directory
	files, err := os.ReadDir(specDir)
	if err != nil {
		panic(err)
	}

	modelRegistry := types.NewRegistry(goPkgRoot, kotlinPkgRoot)

	actorRoleObj := types.Object{
		Name: "ActorRole",
		Fields: []types.Field{
			{
				Name: "role",
				Type: "Role",
			},
		},
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Read only YAML files
		filename := file.Name()
		if !isYAMLFile(filename) {
			continue
		}

		// read the file
		fmt.Println("Reading file for actor check", filename)
		data, err := os.ReadFile(filepath.Join(specDir, filename))
		if err != nil {
			panic(err)
		}
		// parse the file
		var objDef ModelFile
		err = yaml.Unmarshal(data, &objDef)
		if err != nil {
			panic(err)
		}

		for _, o := range objDef.Objects {
			if err = modelRegistry.RegisterIfActor(o); err != nil {
				panic(err)
			}
		}
	}

	modelTypes := make(map[string]struct{})

	var eventPermissions *types.ObjectPermissions

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Read only YAML files
		filename := file.Name()
		if !isYAMLFile(filename) {
			continue
		}

		// read the file
		fmt.Println("Reading file", filename)
		data, err := os.ReadFile(filepath.Join(specDir, filename))
		if err != nil {
			panic(err)
		}
		// parse the file
		var objDef ModelFile
		err = yaml.Unmarshal(data, &objDef)
		if err != nil {
			panic(err)
		}

		// register the permissions
		err = modelRegistry.RegisterPermissions(&actorRoleObj, objDef.Permissions)
		if err != nil {
			panic(err)
		}

		// register the enums
		for _, e := range objDef.Enums {
			err := modelRegistry.RegisterEnum(e)
			if err != nil {
				panic(err)
			}
		}

		// register the objects
		for _, o := range objDef.Objects {
			err := modelRegistry.RegisterObject(o)
			if err != nil {
				panic(err)
			}
			if o.HasCollection() {
				modelTypes[o.Name] = struct{}{}
			}
		}

		// register the errors
		for _, e := range objDef.Errors {
			err := modelRegistry.RegisterError(e)
			if err != nil {
				panic(err)
			}
		}

		// register the events
		for _, e := range objDef.Events {
			err := modelRegistry.RegisterEvent(e)
			if err != nil {
				panic(err)
			}
		}

		// check for event permissions
		if len(objDef.EventPermissions.Read) != 0 {
			if eventPermissions != nil {
				panic(fmt.Errorf("multiple event permissions found"))
			}
			eventPermissions = &objDef.EventPermissions
		}
	}

	// register an enum for the model types
	enum := types.Enum{
		Name:   "Model",
		Type:   types.FieldTypeString,
		Values: make([]string, 0, len(modelTypes)),
	}
	for modelType := range modelTypes {
		enum.Values = append(enum.Values, modelType)
		sort.Strings(enum.Values)
	}
	err = modelRegistry.RegisterEnum(enum)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf("failed to register enum %s", enum.Name)))
	}

	// register the event enums
	for _, eventEnum := range modelRegistry.GetEventEnums() {
		err := modelRegistry.RegisterEnum(eventEnum)
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to register event enum %s", eventEnum.Name)))
		}
	}

	// register the event objects
	for _, eventObj := range modelRegistry.GetEventObjects(eventPermissions) {
		err := modelRegistry.RegisterObject(eventObj)
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to register event object %s", eventObj.Name)))
		}
	}

	// register the actorRole object
	if err := modelRegistry.RegisterObject(actorRoleObj); err != nil {
		panic(err)
	}

	err = modelRegistry.BuildRefs()
	if err != nil {
		panic(errors.Join(err, fmt.Errorf("failed to build refs")))
	}

	err = modelRegistry.ApplyExtendedRolePermissions()
	if err != nil {
		panic(errors.Join(err, fmt.Errorf("failed to apply extended role permissions")))
	}

	fmt.Println("validating spec...")
	err = modelRegistry.Validate()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully validated spec")

	fmt.Println("generating models...")

	var outFiles []templates.OutFile

	if shouldGenGo {
		if goOutDir == "" {
			panic(fmt.Errorf("goOutDir is required"))
		}
		if goPkgRoot == "" {
			panic(fmt.Errorf("goPkgRoot is required"))
		}
		goFiles, err := model_template_go.Gen(model_template_go.GenParams{
			Registry: modelRegistry,
			OutDir:   goOutDir,
			PkgRoot:  goPkgRoot,
		})
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to generate go files")))
		}
		outFiles = append(outFiles, goFiles...)
	}

	if shouldGenTS {
		tsFiles, err := model_template_ts.Gen(model_template_ts.GenParams{
			Registry: modelRegistry,
			OutDir:   tsOutDir,
		})
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to generate ts files")))
		}
		outFiles = append(outFiles, tsFiles...)
	}

	if shouldGenKotlin {
		kotlinFiles, err := model_template_kotlin.Gen(model_template_kotlin.GenParams{
			Registry: modelRegistry,
			OutDir:   kotlinOutDir,
			PkgRoot:  kotlinPkgRoot,
		})
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to generate kotlin files")))
		}
		outFiles = append(outFiles, kotlinFiles...)
	}

	for _, f := range outFiles {
		fileDir := filepath.Dir(f.Path)
		err := os.MkdirAll(fileDir, 0755)
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to create directory %s", fileDir)))
		}
		err = os.WriteFile(f.Path, f.Data, 0644)
		if err != nil {
			panic(errors.Join(err, fmt.Errorf("failed to write file %s", f.Path)))
		}
	}
}

func sanitizePath(path string) string {
	if path != "" && path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}

// isYAMLFile checks if a filename has a YAML extension
func isYAMLFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".yaml" || ext == ".yml"
}
