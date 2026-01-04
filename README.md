# Forge

A comprehensive code generation system that transforms YAML specifications into production-ready full-stack applications with Go, TypeScript, and Kotlin support. This tool automatically generates models, API endpoints, validation logic, permission systems, HTTP handlers, database operations, and React integration—eliminating boilerplate and ensuring consistency across your entire stack.

**Supported Databases:** MongoDB (PostgreSQL and SQLite coming soon)

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Documentation](#documentation)
- [Example](#example)

## Overview

Forge is a sophisticated code generation tool that transforms declarative YAML specifications into a complete, type-safe, full-stack application. It automates the creation of:

### Backend (Go)

- **MongoDB Models**: Complete type-safe models with conversion utilities
- **Database Layer**: CRUD operations, search, lookups, aggregations
- **HTTP Handlers**: RESTful API endpoints with automatic routing
- **Permission System**: RBAC and ABAC enforcement at API level
- **Custom Hooks**: Extensible API behavior via hooks
- **Validation**: Field-level validation with detailed error handling
- **Index Management**: Automatic MongoDB index creation

### Frontend (TypeScript)

- **Type Definitions**: Complete TypeScript interfaces matching backend models
- **API Clients**: Type-safe API functions with automatic serialization
- **React Query Hooks**: TanStack Query integration for data fetching
- **Form State Management**: Pre-built form state with validation
- **MUI Components**: Material-UI DataGrid configurations
- **Permission Utilities**: Client-side permission checking

### Mobile (Kotlin)

- **Data Models**: Kotlin data classes with serialization
- **API Integration**: Type-safe HTTP clients
- **Permission Logic**: Mobile-compatible permission checking

### Key Benefits

- **Single Source of Truth**: YAML specs define your entire data model
- **Type Safety**: End-to-end type checking across all layers
- **Zero Boilerplate**: Automatic generation of repetitive code
- **Consistency**: Guaranteed alignment between frontend, backend, and mobile
- **Permission-First**: Built-in RBAC/ABAC throughout the stack
- **Production Ready**: Generated code includes error handling, validation, and best practices

## Quick Start

### 1. Install Forge

```bash
# Install latest version
go install github.com/JacobDoucet/forge@latest

# Or install a specific version
go install github.com/JacobDoucet/forge@v1.0.0
```

Or build from source:

```bash
git clone https://github.com/JacobDoucet/forge.git
cd forge
task build   # or: go build -o forge .
./forge version
```

### 2. Initialize a New Project

```bash
forge init
```

This creates a `./models` directory with an example YAML model.

### 3. Generate Code

```bash
forge build \
  --specDir=./models \
  --goOutDir=./generated \
  --goPkgRoot=myapp/generated
```

### 4. Use Generated Code

**Go Backend:**

```go
import "myapp/generated/product"

// Create a product
p := product.Model{
    Name:        "Widget",
    Price:       29.99,
    Description: "A useful widget",
}
```

**TypeScript Frontend:**

```typescript
import { Product } from "@/lib/model/product-model";
import { useProductQueries } from "@/lib/react/tanstack-query/product-queries";

// Fetch products with React Query
const { data: products } = useProductQueries().useSearchProducts({});
```

## Documentation

| Document                                            | Description                                           |
| --------------------------------------------------- | ----------------------------------------------------- |
| [CLI Reference](docs/cli-reference.md)              | Complete command-line interface documentation         |
| [YAML Specification](docs/yaml-specification.md)    | Full guide to defining objects, fields, and permissions |
| [Generated Code](docs/generated-code.md)            | Structure and usage of generated Go, TypeScript, and Kotlin code |
| [Advanced Features](docs/advanced-features.md)      | Object references, parent-child relationships, custom ABAC |
| [Examples](docs/examples.md)                        | Complete real-world examples                          |
| [Development Guide](docs/development.md)            | Extending Forge with new types and templates          |
| [Architecture](docs/architecture.md)                | How Forge works internally and troubleshooting        |

## Example

The `example/` directory contains a complete example for local development:

```bash
cd example
./build.sh
```

This will:

1. Build the forge CLI
2. Generate Go code from example models
3. Verify the generated code compiles

See [example/README.md](example/README.md) for more details.

---

**Generated Code Notice**: All generated files include a header comment:

```
// This file is auto-generated. DO NOT EDIT.
```

Never manually edit generated files. Instead, modify the YAML specifications and regenerate.
