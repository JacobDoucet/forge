# CLI Reference

[← Back to README](../README.md)

## forge init

Initialize a new Forge project with example models.

```bash
forge init [flags]
```

**Flags:**

| Flag    | Short | Default    | Description                   |
| ------- | ----- | ---------- | ----------------------------- |
| `--dir` | `-d`  | `./models` | Directory to create models in |

**Example:**

```bash
# Initialize with default directory
forge init

# Initialize with custom directory
forge init --dir ./my-models
```

## forge build

Generate code from YAML model specifications.

```bash
forge build [flags]
```

**Flags:**

| Flag              | Required       | Description                           | Example              |
| ----------------- | -------------- | ------------------------------------- | -------------------- |
| `--specDir`       | Yes            | Directory containing YAML specs       | `./models`           |
| `--goOutDir`      | For Go         | Output directory for Go files         | `./generated`        |
| `--goPkgRoot`     | For Go         | Go package root path                  | `myapp/generated`    |
| `--tsOutDir`      | For TypeScript | Output directory for TypeScript files | `./frontend/lib`     |
| `--kotlinOutDir`  | For Kotlin     | Output directory for Kotlin files     | `./mobile/src`       |
| `--kotlinPkgRoot` | For Kotlin     | Kotlin package root                   | `com.example.models` |

**Examples:**

```bash
# Generate Go code
forge build \
  --specDir=./models \
  --goOutDir=./generated \
  --goPkgRoot=myapp/generated

# Generate TypeScript code
forge build \
  --specDir=./models \
  --tsOutDir=./frontend/lib

# Generate all target languages
forge build \
  --specDir=./models \
  --goOutDir=./backend/models \
  --goPkgRoot=myapp/models \
  --tsOutDir=./frontend/lib \
  --kotlinOutDir=./mobile/src \
  --kotlinPkgRoot=com.example.models
```

## forge version

Display the current version of Forge.

```bash
forge version
```

## Configuration File (.forge.yml)

Instead of passing flags on every build, you can create a `.forge.yml` configuration file:

```yaml
# Forge Configuration

# Required version of forge CLI (optional)
# If set, forge will check the version and prompt to update if it doesn't match
forgeVersion: "1.2.0"

# Directory containing YAML spec files
specDir: ./models

# Go output configuration
go:
  outDir: ./generated
  pkgRoot: myapp/generated

# TypeScript output configuration (optional)
typescript:
  outDir: ./frontend/lib

# Kotlin output configuration (optional)
kotlin:
  outDir: ./mobile/src
  pkgRoot: com.example.models
```

### Version Checking

When `forgeVersion` is specified in your config file, Forge will:

1. Compare the current CLI version with the required version
2. If they don't match, prompt you to update
3. If you accept, automatically run `go install` to update to the specified version

**Example interaction:**

```
⚠️  Version mismatch detected!
   Current version:  1.1.0
   Required version: 1.2.0

Would you like to update forge to version 1.2.0? [y/N]: y
Running: go install github.com/JacobDoucet/forge@v1.2.0

✅ Successfully updated to version 1.2.0
Please re-run your command.
```

This ensures all team members use the same version of Forge, preventing inconsistencies in generated code.
