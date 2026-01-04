package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"d3tech.com/platform/templates"
	model_template_go "d3tech.com/platform/templates/go"
	model_template_kotlin "d3tech.com/platform/templates/kotlin"
	model_template_ts "d3tech.com/platform/templates/ts"
	"d3tech.com/platform/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate code from YAML model specifications",
	Long: `Build generates Go, TypeScript, and/or Kotlin code from YAML model specifications.
Specify at least one output directory to generate code for that language.`,
	RunE: runBuild,
}

var (
	goOutDir      string
	goPkgRoot     string
	tsOutDir      string
	kotlinOutDir  string
	kotlinPkgRoot string
	specDir       string
)

func init() {
	buildCmd.Flags().StringVar(&goOutDir, "goOutDir", "", "output directory for generated Go files")
	buildCmd.Flags().StringVar(&goPkgRoot, "goPkgRoot", "", "root package for generated Go files")
	buildCmd.Flags().StringVar(&tsOutDir, "tsOutDir", "", "output directory for generated TypeScript files")
	buildCmd.Flags().StringVar(&kotlinOutDir, "kotlinOutDir", "", "output directory for generated Kotlin files")
	buildCmd.Flags().StringVar(&kotlinPkgRoot, "kotlinPkgRoot", "", "root package for generated Kotlin files")
	buildCmd.Flags().StringVar(&specDir, "specDir", "", "directory containing YAML spec files")
	buildCmd.MarkFlagRequired("specDir")
}

type ModelFile struct {
	Objects          []types.Object          `yaml:"objects"`
	Enums            []types.Enum            `yaml:"enums"`
	Permissions      types.PermissionsDef    `yaml:"permissions"`
	Errors           []types.CustomError     `yaml:"errors"`
	Events           []string                `yaml:"events"`
	EventPermissions types.ObjectPermissions `yaml:"eventPermissions"`
}

func runBuild(cmd *cobra.Command, args []string) error {
	shouldGenGo := goOutDir != "" || goPkgRoot != ""
	shouldGenTS := tsOutDir != ""
	shouldGenKotlin := kotlinOutDir != "" || kotlinPkgRoot != ""

	if !shouldGenGo && !shouldGenTS && !shouldGenKotlin {
		return fmt.Errorf("at least one output must be specified (--goOutDir, --tsOutDir, or --kotlinOutDir)")
	}

	goOutDir = sanitizePath(goOutDir)
	goPkgRoot = sanitizePath(goPkgRoot)
	tsOutDir = sanitizePath(tsOutDir)
	kotlinOutDir = sanitizePath(kotlinOutDir)

	// ensure out dir
	if goOutDir != "" {
		if err := os.MkdirAll(goOutDir, 0755); err != nil {
			return fmt.Errorf("failed to create Go output directory: %w", err)
		}
	}

	// Check if spec directory exists
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		return fmt.Errorf("spec directory %s does not exist", specDir)
	}

	// Read files from the spec directory
	files, err := os.ReadDir(specDir)
	if err != nil {
		return fmt.Errorf("failed to read spec directory: %w", err)
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
			continue
		}

		filename := file.Name()
		if !isYAMLFile(filename) {
			continue
		}

		fmt.Println("Reading file for actor check", filename)
		data, err := os.ReadFile(filepath.Join(specDir, filename))
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filename, err)
		}

		var objDef ModelFile
		if err = yaml.Unmarshal(data, &objDef); err != nil {
			return fmt.Errorf("failed to parse YAML file %s: %w", filename, err)
		}

		for _, o := range objDef.Objects {
			if err = modelRegistry.RegisterIfActor(o); err != nil {
				return fmt.Errorf("failed to register actor from %s: %w", filename, err)
			}
		}
	}

	modelTypes := make(map[string]struct{})
	var eventPermissions *types.ObjectPermissions

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if !isYAMLFile(filename) {
			continue
		}

		fmt.Println("Reading file", filename)
		data, err := os.ReadFile(filepath.Join(specDir, filename))
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filename, err)
		}

		var objDef ModelFile
		if err = yaml.Unmarshal(data, &objDef); err != nil {
			return fmt.Errorf("failed to parse YAML file %s: %w", filename, err)
		}

		if err = modelRegistry.RegisterPermissions(&actorRoleObj, objDef.Permissions); err != nil {
			return fmt.Errorf("failed to register permissions from %s: %w", filename, err)
		}

		for _, e := range objDef.Enums {
			if err := modelRegistry.RegisterEnum(e); err != nil {
				return fmt.Errorf("failed to register enum %s from %s: %w", e.Name, filename, err)
			}
		}

		for _, o := range objDef.Objects {
			if err := modelRegistry.RegisterObject(o); err != nil {
				return fmt.Errorf("failed to register object %s from %s: %w", o.Name, filename, err)
			}
			if o.HasCollection() {
				modelTypes[o.Name] = struct{}{}
			}
		}

		for _, e := range objDef.Errors {
			if err := modelRegistry.RegisterError(e); err != nil {
				return fmt.Errorf("failed to register error %s from %s: %w", e.Name, filename, err)
			}
		}

		for _, e := range objDef.Events {
			if err := modelRegistry.RegisterEvent(e); err != nil {
				return fmt.Errorf("failed to register event %s from %s: %w", e, filename, err)
			}
		}

		if len(objDef.EventPermissions.Read) != 0 {
			if eventPermissions != nil {
				return fmt.Errorf("multiple event permissions found")
			}
			eventPermissions = &objDef.EventPermissions
		}
	}

	// Register an enum for the model types
	enum := types.Enum{
		Name:   "Model",
		Type:   types.FieldTypeString,
		Values: make([]string, 0, len(modelTypes)),
	}
	for modelType := range modelTypes {
		enum.Values = append(enum.Values, modelType)
		sort.Strings(enum.Values)
	}
	if err = modelRegistry.RegisterEnum(enum); err != nil {
		return errors.Join(err, fmt.Errorf("failed to register enum %s", enum.Name))
	}

	// Register the event enums
	for _, eventEnum := range modelRegistry.GetEventEnums() {
		if err := modelRegistry.RegisterEnum(eventEnum); err != nil {
			return errors.Join(err, fmt.Errorf("failed to register event enum %s", eventEnum.Name))
		}
	}

	// Register the event objects
	for _, eventObj := range modelRegistry.GetEventObjects(eventPermissions) {
		if err := modelRegistry.RegisterObject(eventObj); err != nil {
			return errors.Join(err, fmt.Errorf("failed to register event object %s", eventObj.Name))
		}
	}

	// Register the actorRole object
	if err := modelRegistry.RegisterObject(actorRoleObj); err != nil {
		return fmt.Errorf("failed to register ActorRole object: %w", err)
	}

	if err = modelRegistry.BuildRefs(); err != nil {
		return errors.Join(err, fmt.Errorf("failed to build refs"))
	}

	if err = modelRegistry.ApplyExtendedRolePermissions(); err != nil {
		return errors.Join(err, fmt.Errorf("failed to apply extended role permissions"))
	}

	fmt.Println("Validating spec...")
	if err = modelRegistry.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	fmt.Println("Successfully validated spec")

	fmt.Println("Generating models...")

	var outFiles []templates.OutFile

	if shouldGenGo {
		if goOutDir == "" {
			return fmt.Errorf("goOutDir is required when generating Go code")
		}
		if goPkgRoot == "" {
			return fmt.Errorf("goPkgRoot is required when generating Go code")
		}
		goFiles, err := model_template_go.Gen(model_template_go.GenParams{
			Registry: modelRegistry,
			OutDir:   goOutDir,
			PkgRoot:  goPkgRoot,
		})
		if err != nil {
			return errors.Join(err, fmt.Errorf("failed to generate Go files"))
		}
		outFiles = append(outFiles, goFiles...)
	}

	if shouldGenTS {
		tsFiles, err := model_template_ts.Gen(model_template_ts.GenParams{
			Registry: modelRegistry,
			OutDir:   tsOutDir,
		})
		if err != nil {
			return errors.Join(err, fmt.Errorf("failed to generate TypeScript files"))
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
			return errors.Join(err, fmt.Errorf("failed to generate Kotlin files"))
		}
		outFiles = append(outFiles, kotlinFiles...)
	}

	for _, f := range outFiles {
		fileDir := filepath.Dir(f.Path)
		if err := os.MkdirAll(fileDir, 0755); err != nil {
			return errors.Join(err, fmt.Errorf("failed to create directory %s", fileDir))
		}
		if err := os.WriteFile(f.Path, f.Data, 0644); err != nil {
			return errors.Join(err, fmt.Errorf("failed to write file %s", f.Path))
		}
	}

	fmt.Printf("✓ Generated %d files\n", len(outFiles))
	return nil
}

func sanitizePath(path string) string {
	if path != "" && path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}

func isYAMLFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".yaml" || ext == ".yml"
}
