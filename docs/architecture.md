# Architecture & Troubleshooting

[← Back to README](../README.md)

This guide explains how Forge works internally and provides solutions for common issues.

## Table of Contents

- [Architecture](#architecture)
  - [Code Generation Pipeline](#code-generation-pipeline)
  - [Registry System](#registry-system)
  - [Template Execution](#template-execution)
  - [Multi-Language Support](#multi-language-support)
- [Troubleshooting](#troubleshooting)
  - [Common Issues](#common-issues)
  - [Validation Checklist](#validation-checklist)

## Architecture

### Code Generation Pipeline

```
┌─────────────────┐
│  YAML Specs     │
│  (*.yaml)       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Parse YAML     │
│  (yaml.v3)      │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Build Registry │
│  - Objects      │
│  - Enums        │
│  - Permissions  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Validate Specs │
│  - Types        │
│  - References   │
│  - Permissions  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Execute        │
│  Templates      │
│  - Go           │
│  - TypeScript   │
│  - Kotlin       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Write Files    │
│  to Output Dirs │
└─────────────────┘
```

### Registry System

The `types.Registry` maintains the global state:

```go
type Registry struct {
    objects       map[string]*Object
    enums         map[string]*Enum
    permissions   *PermissionsDef
    errors        []CustomError
    events        []string
}
```

**Key operations:**

- `RegisterObject()` - Add objects and validate
- `RegisterEnum()` - Add enums
- `BuildRefs()` - Resolve references between objects
- `Validate()` - Comprehensive validation

### Template Execution

Templates use Go's `text/template` with custom functions:

```go
tmpl := template.New("model").Funcs(template.FuncMap{
    "ToLower":    strings.ToLower,
    "ToCamel":    utils.ToCamel,
    "GoType":     utils.GoType,
    "TSType":     utils.TSType,
    // ... more functions
})

tmpl.Parse(templateString)
tmpl.Execute(writer, data)
```

### Multi-Language Support

Each language has its own generator package:

- `templates/go/gen.go` - Go backend
- `templates/ts/gen.go` - TypeScript frontend
- `templates/kotlin/gen.go` - Kotlin mobile

All share the same `types.Registry` for consistency.

## Troubleshooting

### Common Issues

#### 1. "Unknown field type"

**Error:**

```
panic: unknown field type: MyType
```

**Solution:**

- Check that custom types are defined as objects or enums
- Verify type name spelling matches exactly (case-sensitive)
- For references, use `Ref<Type>` or `ParentRef<Type>`

#### 2. "Circular reference detected"

**Error:**

```
panic: circular reference: User -> Account -> User
```

**Solution:**

- Use `Ref<Type>` for one direction instead of embedding
- Review object relationships for cycles
- Consider using IDs instead of embedded objects

#### 3. "Permission role not found"

**Error:**

```
panic: permission role AdminRole not found in registry
```

**Solution:**

- Define all roles in `permissions.yaml`
- Ensure `permissions.yaml` is in the spec directory
- Check role name spelling (case-sensitive)

#### 4. "Index field does not exist"

**Error:**

```
panic: index field "createdAt" does not exist on object User
```

**Solution:**

- Use correct field names (check case)
- For nested fields, use dot notation: `created.at`
- Verify the field is actually defined on the object

#### 5. Generated Go code won't compile

**Symptoms:**

- Import errors
- Type mismatch errors

**Solutions:**

- Run `go mod tidy` in backend directory
- Check `--goPkgRoot` matches your actual package path
- Verify all dependencies are in `go.mod`

#### 6. TypeScript errors in generated code

**Symptoms:**

- "Cannot find module" errors
- Type errors

**Solutions:**

- Check `--tsOutDir` path is correct
- Verify TypeScript compiler settings in `tsconfig.json`
- Run `npm install` to ensure dependencies are present

### Validation Checklist

Before running the generator:

- [ ] All YAML files are valid YAML syntax
- [ ] All referenced types are defined (objects, enums)
- [ ] No circular references
- [ ] All permission roles defined in `permissions.yaml`
- [ ] All ABAC rules have corresponding field definitions
- [ ] Index fields exist on their objects
- [ ] Unique indexes are on appropriate fields
- [ ] Actor configuration uses existing fields

### Getting Help

If you encounter issues:

1. **Check validation output** - The generator provides detailed error messages
2. **Review examples** - Look at existing model definitions in `example/models/`
3. **Inspect generated code** - Generated files include comments and structure
4. **Check types** - Ensure field types match available types

---

**Generated Code Notice**: All generated files include a header comment:

```
// This file is auto-generated. DO NOT EDIT.
```

Never manually edit generated files. Instead, modify the YAML specifications and regenerate.
