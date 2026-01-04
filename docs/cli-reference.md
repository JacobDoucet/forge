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
