# Example Module

This directory contains an example module for local development and testing of Forge.

## Structure

```
example/
├── models/                  # YAML model specifications
│   └── example.yaml         # Example model definition
├── generated/               # Generated code (created by build.sh)
│   ├── go/                  # Go backend code
│   ├── typescript/          # TypeScript frontend code
│   └── kotlin/              # Kotlin mobile code
├── build.sh                 # Build script
├── go.mod                   # Go module definition
└── README.md                # This file
```

## Usage

### Generate Code

Run the build script to generate code from the example models:

```bash
./build.sh
```

Or use the Taskfile from the repo root:

```bash
task example:build
```

This will:

1. Build the forge CLI tool
2. Generate Go, TypeScript, and Kotlin code from `models/example.yaml`
3. Output the generated files to `generated/{go,typescript,kotlin}/`
4. Verify the Go generated code compiles

### Manual Build

You can also run the steps manually:

```bash
# From the forge root directory
go build -o forge .

# Generate code
./forge build \
    --specDir ./example/models \
    --goOutDir ./example/generated/go \
    --goPkgRoot github.com/JacobDoucet/forge/example/generated/go \
    --tsOutDir ./example/generated/typescript \
    --kotlinOutDir ./example/generated/kotlin \
    --kotlinPkgRoot com.forge.example.generated
```

## Example Model

The example model (`models/example.yaml`) demonstrates:

- **Enums**: `TaskStatus`, `TaskPriority`
- **Objects with Collections**: `Task`, `Project` (MongoDB-backed)
- **Nested Objects**: `TaskComment` (embedded type)
- **HTTP Endpoints**: REST API configuration
- **Permissions**: RBAC configuration
- **Indexes**: MongoDB index definitions
- **Errors**: Custom error types
- **Events**: Event type definitions

## Cleaning Up

To remove generated files:

```bash
rm -rf generated/
```
