# Development Guide

[← Back to README](../README.md)

This guide covers extending Forge with new field types, custom templates, and permission system extensions.

## Table of Contents

- [Adding New Field Types](#adding-new-field-types)
- [Creating Custom Templates](#creating-custom-templates)
- [Extending the Permission System](#extending-the-permission-system)
- [Running Tests](#running-tests)
- [Debugging Generated Code](#debugging-generated-code)

## Adding New Field Types

To support a new field type across the entire stack:

### 1. Add type to parser

In `types/field.go`:

```go
const (
    FieldTypeString    = "string"
    FieldTypeInt       = "int"
    // Add your new type
    FieldTypeMyNewType = "mynewtype"
)
```

### 2. Add Go type mapping

In `templates/go/utils.go`:

```go
func GoType(field types.Field) string {
    switch field.Type {
    case types.FieldTypeMyNewType:
        return "MyNewGoType"
    // ...
    }
}
```

### 3. Add TypeScript type mapping

In `templates/ts/utils.go`:

```go
func TSType(field types.Field) string {
    switch field.Type {
    case types.FieldTypeMyNewType:
        return "MyNewTSType"
    // ...
    }
}
```

### 4. Update MongoDB conversion

In model templates, add conversion logic:

- Add conversion logic in `ToMongoRecord`
- Add parsing logic in `FromMongoRecord`

### 5. Update validation

In `templates/go/utils_template_model.go`:

```go
// Add validation rules for your new type
```

## Creating Custom Templates

### 1. Create Template File

Create a new file in `templates/go/` or `templates/ts/`:

```go
// templates/go/template_my_custom.go
package model_template_go

const MyCustomTemplate = `
// Custom generated code for {{ .Object.Name }}
package {{ .PackageName }}

func CustomFunction() {
    // Your template logic
}
`
```

### 2. Register Template

Add to `gen.go`:

```go
func Gen(params GenParams) ([]templates.OutFile, error) {
    // ... existing code

    // Add your custom template
    myCustomFiles, err := genMyCustom(params)
    if err != nil {
        return nil, err
    }
    outFiles = append(outFiles, myCustomFiles...)

    return outFiles, nil
}

func genMyCustom(params GenParams) ([]templates.OutFile, error) {
    // Generate your custom files
    for _, obj := range params.Registry.GetObjects() {
        data := struct {
            Object      types.Object
            PackageName string
        }{
            Object:      obj,
            PackageName: strings.ToLower(obj.Name),
        }

        content, err := executeTemplate(MyCustomTemplate, data)
        if err != nil {
            return nil, err
        }

        outFiles = append(outFiles, templates.OutFile{
            Path: filepath.Join(params.OutDir, obj.Name, "custom.go"),
            Data: []byte(content),
        })
    }

    return outFiles, nil
}
```

### 3. Use Template Functions

Available template functions (defined in `template_funcs.go`):

```go
{{ .Object.Name | ToLower }}           // Convert to lowercase
{{ .Object.Name | ToCamel }}           // Convert to camelCase
{{ .Object.Name | ToSnake }}           // Convert to snake_case
{{ .Field.Type | GoType }}             // Get Go type
{{ .Field.Type | TSType }}             // Get TypeScript type
{{ if .Object.HasCollection }}...{{ end }}  // Conditional
```

## Extending the Permission System

### Add New Permission Type

#### 1. Define in `permissions.yaml`:

```yaml
permissions:
  rbac:
    - name: MyNewRole
      customPermissions:
        - MyCustomPermission
```

#### 2. Update permission templates:

- `template_permissions_actor.go` - Actor interface
- `template_obj__model_permissions.go` - Permission checking
- `template_permissions_obj__can_access.go` - TypeScript utilities

#### 3. Implement custom logic:

```go
// In generated permissions.go
func CanPerformCustomAction(actor permissions.Actor, obj *Model) bool {
    // Your custom permission logic
    if actor.HasPermission("MyCustomPermission") {
        return true
    }
    return false
}
```

## Running Tests

```bash
# Test the generator
go test ./...

# Test generated Go code
cd example
./build.sh
go test ./generated/...

# Test with verbose output
go test -v ./...
```

## Debugging Generated Code

### 1. Enable verbose output

```bash
forge build --specDir=./models --goOutDir=./generated --goPkgRoot=myapp/generated
```

### 2. Check validation errors

The generator validates specs before generation. Common errors:

- Unknown field types
- Circular references
- Invalid permission configurations
- Duplicate object names

### 3. Inspect generated files

- Go files include package imports and type info
- TypeScript files are formatted with Prettier
- All files have generation timestamp comments

### 4. Common debugging patterns

**Check if a field is being processed:**

Add debug output in template functions:

```go
func MyTemplateFunc(field types.Field) string {
    log.Printf("Processing field: %s, type: %s", field.Name, field.Type)
    // ...
}
```

**Validate template output:**

```bash
# Check for Go syntax errors
gofmt -e ./generated/**/*.go

# Check for TypeScript errors
npx tsc --noEmit
```
