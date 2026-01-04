# Example Module

This directory contains an example module for local development and testing of Forge.

## Structure

```
example/
├── models/          # YAML model specifications
│   └── example.yaml # Example model definition
├── generated/       # Generated Go code (created by build.sh)
├── build.sh         # Build script
├── go.mod           # Go module definition
└── README.md        # This file
```

## Usage

### Generate Code

Run the build script to generate Go code from the example models:

```bash
./build.sh
```

This will:

1. Build the forge CLI tool
2. Generate Go code from `models/example.yaml`
3. Output the generated files to `generated/`
4. Verify the generated code compiles

### Manual Build

You can also run the steps manually:

```bash
# From the forge root directory
go build -o forge .

# Generate code
./forge build \
    --specDir ./example/models \
    --goOutDir ./example/generated \
    --goPkgRoot d3tech.com/platform/example/generated
```

## Example Model

The example model (`models/example.yaml`) demonstrates:

- **Enums**: `TaskStatus`, `TaskPriority`, `Role`
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
