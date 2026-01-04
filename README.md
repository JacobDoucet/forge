# Platform Model Generator

A comprehensive code generation system that transforms YAML specifications into production-ready full-stack applications with MongoDB, Go, TypeScript, and Kotlin support. This tool automatically generates models, API endpoints, validation logic, permission systems, HTTP handlers, database operations, and React integration—eliminating boilerplate and ensuring consistency across your entire stack.

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [YAML Specification Guide](#yaml-specification-guide)
  - [Object Definition](#object-definition)
  - [Field Types Reference](#field-types-reference)
  - [Collections](#collections)
  - [Indexes](#indexes)
  - [HTTP Endpoints](#http-endpoints)
  - [Permissions System](#permissions-system)
  - [Field Validation](#field-validation)
  - [Actors and Authentication](#actors-and-authentication)
  - [Enums](#enums)
  - [Errors and Events](#errors-and-events)
- [Generated Code Reference](#generated-code-reference)
  - [Go Backend Structure](#go-backend-structure)
  - [TypeScript Frontend Structure](#typescript-frontend-structure)
  - [Kotlin Mobile Structure](#kotlin-mobile-structure)
- [Advanced Features](#advanced-features)
  - [Object References](#object-references)
  - [Parent-Child Relationships](#parent-child-relationships)
  - [Custom ABAC Rules](#custom-abac-rules)
  - [Field-Level Permissions](#field-level-permissions)
  - [Immutable Fields](#immutable-fields)
- [Practical Examples](#practical-examples)
- [Development Guide](#development-guide)
- [Architecture & How It Works](#architecture--how-it-works)
- [Troubleshooting](#troubleshooting)

## Overview

The Platform Model Generator is a sophisticated code generation tool that transforms declarative YAML specifications into a complete, type-safe, full-stack application. It automates the creation of:

### Backend (Go)

- **MongoDB Models**: Complete type-safe models with conversion utilities
- **Database Layer**: CRUD operations, search, lookups, aggregations
- **HTTP Handlers**: RESTful API endpoints with automatic routing
- **Permission System**: RBAC and ABAC enforcement at API level
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

1. **Define Your Model** - Create a YAML file in `../platform_models/`:

```yaml
# model_product.yaml
objects:
  - name: Product
    fields:
      - name: name
        type: string
        validate:
          required: true
      - name: price
        type: float64
        validate:
          required: true
          min: 0
      - name: description
        type: string
    collection:
      - type: mongo
        name: products
    http:
      endpoint: products
      methods:
        - GET
        - POST
        - PATCH
        - DELETE
    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
      write:
        - rbac: PlatformAdmin
```

2. **Generate Code** - Run the generator:

```bash
go run main.go \
  --goOutDir=../backend/models \
  --goPkgRoot=ns-vpn.com/api/models/ \
  --tsOutDir=../frontend/lib \
  --specDir=../platform_models
```

3. **Use Generated Code**:

**Go Backend:**

```go
import "ns-vpn.com/api/models/product"

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

## Usage

### Command Line Interface

```bash
go run main.go [flags]
```

### Flags

| Flag              | Required       | Description                           | Example                  |
| ----------------- | -------------- | ------------------------------------- | ------------------------ |
| `--goOutDir`      | For Go         | Output directory for Go files         | `../backend/models`      |
| `--goPkgRoot`     | For Go         | Go package root path                  | `ns-vpn.com/api/models/` |
| `--tsOutDir`      | For TypeScript | Output directory for TypeScript files | `../frontend/lib`        |
| `--kotlinOutDir`  | For Kotlin     | Output directory for Kotlin files     | `../mobile/src`          |
| `--kotlinPkgRoot` | For Kotlin     | Kotlin package root                   | `com.example.models`     |
| `--specDir`       | Yes            | Directory containing YAML specs       | `../platform_models`     |

### Full Example

```bash
# Generate all target languages
go run main.go \
  --goOutDir=../backend/models \
  --goPkgRoot=ns-vpn.com/api/models/ \
  --tsOutDir=../frontend/lib \
  --kotlinOutDir=../mobile/app/src/main/java/com/example/models \
  --kotlinPkgRoot=com.example.models \
  --specDir=../platform_models

# Generate only Go
go run main.go \
  --goOutDir=../backend/models \
  --goPkgRoot=ns-vpn.com/api/models/ \
  --specDir=../platform_models

# Generate only TypeScript
go run main.go \
  --tsOutDir=../frontend/lib \
  --specDir=../platform_models
```

### Workflow

1. **Create/Edit YAML specs** in `platform_models/` directory
2. **Run the generator** with appropriate flags
3. **Generated files are written** to specified output directories
4. **Import and use** the generated code in your application

The generator:

- Validates all YAML specifications before generating
- Creates directory structures automatically
- Overwrites existing generated files (marked with `// This file is auto-generated. DO NOT EDIT.`)
- Preserves custom code in non-generated files

## YAML Specification Guide

### Object Definition

Objects are the core building blocks. Each object definition creates a complete set of model, API, database, and UI code.

#### Basic Structure

```yaml
objects:
  - name: User # Object name (PascalCase)
    fields: [] # Field definitions
    collection: [] # Database configuration
    http: {} # HTTP endpoint configuration
    permissions: {} # Access control rules
    abac: [] # Custom ABAC rules
    indexes: [] # Database indexes
    actor: {} # Actor configuration (for authentication)
```

#### Complete Example

```yaml
objects:
  - name: User
    # Fields define the data structure
    fields:
      - name: accountId
        type: ParentRef<Account>
        immutable: true # Cannot be changed after creation
      - name: email
        type: string
        validate:
          required: true
          email: true
      - name: firstName
        type: string
      - name: lastName
        type: string
      - name: roles
        type: List<ActorRole>

    # Collection configuration
    collection:
      - type: mongo
        name: users

    # HTTP endpoints
    http:
      endpoint: users # URL path segment
      methods:
        - GET # Get by ID
        - POST # Create
        - PATCH # Update
        - DELETE # Delete

    # Permissions
    permissions:
      read:
        - rbac: PlatformAdmin # Admin can read all
        - rbac: AccountOwner
          abac:
            - AccountId # Account owner can read their account's users
        - abac:
            - ActorId # Users can read themselves
      write:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId

    # Custom ABAC field mappings
    abac:
      - name: AccountId
        field: accountId

    # Database indexes
    indexes:
      - name: email
        fields:
          - name: email
      - name: accountEmail
        unique: true
        fields:
          - name: accountId
          - name: email

    # Actor configuration (for authentication)
    actor:
      id: id
      username: "{email}"
      name: "{firstName} {lastName}"
```

### Field Types Reference

#### Primitive Types

| Type        | Go Type     | TypeScript Type | Kotlin Type | Description                 |
| ----------- | ----------- | --------------- | ----------- | --------------------------- |
| `string`    | `string`    | `string`        | `String`    | Text data                   |
| `int`       | `int`       | `number`        | `Int`       | Integer (platform-specific) |
| `int32`     | `int32`     | `number`        | `Int`       | 32-bit integer              |
| `int64`     | `int64`     | `number`        | `Long`      | 64-bit integer              |
| `float`     | `float64`   | `number`        | `Double`    | Floating point              |
| `float32`   | `float32`   | `number`        | `Float`     | 32-bit float                |
| `float64`   | `float64`   | `number`        | `Double`    | 64-bit float                |
| `bool`      | `bool`      | `boolean`       | `Boolean`   | True/false                  |
| `timestamp` | `time.Time` | `Date`          | `Date`      | Date/time                   |

#### Reference Types

| Type Syntax             | Description                     | Example              |
| ----------------------- | ------------------------------- | -------------------- |
| `Ref<ObjectName>`       | Reference to another object     | `Ref<User>`          |
| `ParentRef<ObjectName>` | Parent reference (hierarchical) | `ParentRef<Account>` |

**Example:**

```yaml
fields:
  - name: ownerId
    type: Ref<User> # Foreign key reference
  - name: accountId
    type: ParentRef<Account> # Parent in hierarchy
```

**Generated behavior:**

- References generate ObjectID fields in MongoDB
- Automatic lookup/aggregation queries
- Foreign key validation

#### Collection Types

| Type Syntax              | Description     | Example           |
| ------------------------ | --------------- | ----------------- |
| `List<Type>`             | Array of values | `List<string>`    |
| `Map<KeyType,ValueType>` | Key-value pairs | `Map<string,int>` |

**Example:**

```yaml
fields:
  - name: tags
    type: List<string>
  - name: metadata
    type: Map<string,string>
  - name: members
    type: List<Ref<User>> # Array of references
```

#### Special Types

| Type         | Description                       | Auto-Generated Fields                                    |
| ------------ | --------------------------------- | -------------------------------------------------------- |
| `ActorTrace` | Tracks who/when for create/update | `at` (timestamp), `by` (actor ID), `byName` (actor name) |

**Example:**

```yaml
fields:
  - name: created
    type: ActorTrace # Auto-populated on creation
  - name: updated
    type: ActorTrace # Auto-updated on modification
```

**Generated structure:**

```typescript
type ActorTrace = {
  at?: Date;
  by?: string;
  byName?: string;
};
```

#### Enum Types

Reference enums defined in the `enums` section:

```yaml
enums:
  - name: Role
    type: string
    values:
      - Admin
      - User
      - Guest

objects:
  - name: User
    fields:
      - name: role
        type: Role # Uses the enum above
```

#### Embedded Object Types

Reference other objects defined as embedded types:

```yaml
# obj_billing_info.yaml
objects:
  - name: BillingInfo
    fields:
      - name: email
        type: string
      - name: phone
        type: string
      - name: address
        type: string

# model_user.yaml
objects:
  - name: User
    fields:
      - name: billingInfo
        type: BillingInfo        # Embedded object, not a reference
```

### Collections

Collections define how objects are persisted in the database.

```yaml
collection:
  - type: mongo
    name: collection_name # MongoDB collection name (plural, snake_case recommended)
```

**Example:**

```yaml
objects:
  - name: Product
    collection:
      - type: mongo
        name: products # Stored in "products" collection
```

**What gets generated:**

- `{model}_mongo/collection.go` - Collection management functions
- Index creation utilities
- CRUD operations
- Search and aggregation pipelines

### Indexes

Define MongoDB indexes to optimize query performance.

#### Simple Index

```yaml
indexes:
  - name: email_idx
    fields:
      - name: email
        order: 1 # 1 = ascending, -1 = descending
```

#### Unique Index

```yaml
indexes:
  - name: email_unique
    unique: true
    fields:
      - name: email
```

#### Compound Index

```yaml
indexes:
  - name: account_email
    unique: true
    fields:
      - name: accountId
        order: 1
      - name: email
        order: 1
```

#### Nested Field Index

```yaml
indexes:
  - name: created_at
    fields:
      - name: created.at # Index on nested field
        order: -1 # Descending (newest first)
```

#### Multi-Field Index Example

```yaml
indexes:
  - name: rolesByAccountId
    fields:
      - name: actorRoles.accountId
        order: 1
      - name: actorRoles.role
        order: 1
      - name: actorRoles.groupId
        order: 1
```

**Generated code:**

- Index creation in `{model}_mongo/collection.go`
- Automatic index creation on application startup
- Index management utilities

### HTTP Endpoints

Define which HTTP methods are available for your object.

#### Basic Configuration

```yaml
http:
  endpoint: users # URL segment (defaults to lowercase object name)
  methods:
    - GET # GET /users/:id
    - POST # POST /users
    - PATCH # PATCH /users/:id
    - DELETE # DELETE /users/:id
```

#### Generated Routes

| Method   | Route           | Description               |
| -------- | --------------- | ------------------------- |
| `GET`    | `/endpoint/:id` | Fetch single object by ID |
| `POST`   | `/endpoint`     | Create new object         |
| `PATCH`  | `/endpoint/:id` | Update existing object    |
| `DELETE` | `/endpoint/:id` | Delete object             |

**Additional auto-generated endpoints:**

- `GET /endpoint` - Search/list with query parameters
- `GET /endpoint/:id/lookup` - Fetch with related objects

#### Custom Endpoint Names

```yaml
http:
  endpoint: custom-path # Creates /custom-path instead of /objectname
  methods:
    - GET
    - POST
```

**What gets generated:**

- HTTP handlers in `{model}_http/handlers.go`
- Route registration in `{model}_http/routes.go`
- Request/response serialization
- Automatic permission checking
- Validation before database operations

### Permissions System

The generator supports sophisticated Role-Based Access Control (RBAC) and Attribute-Based Access Control (ABAC).

#### Permission Structure

```yaml
permissions:
  read: [] # Who can read this object
  write: [] # Who can create/update/delete this object
```

#### RBAC (Role-Based Access Control)

Simple role-based permissions:

```yaml
permissions:
  read:
    - rbac: PlatformAdmin # PlatformAdmin role can read
    - rbac: AccountOwner # AccountOwner role can read
  write:
    - rbac: PlatformAdmin # Only PlatformAdmin can write
```

#### ABAC (Attribute-Based Access Control)

Attribute-based rules that check field values:

```yaml
permissions:
  read:
    - abac:
        - ActorId # User can read if they are the subject
  write:
    - rbac: AccountOwner
      abac:
        - AccountId # AccountOwner can write their own account's objects
```

#### Multiple ABAC Rules (AND Logic)

```yaml
permissions:
  read:
    - rbac: GroupAdmin
      abac:
        - AccountId # Must match account ID
        - GroupId # AND must match group ID
```

#### Permission Rule Evaluation

Rules are evaluated with OR logic between top-level entries, AND logic within ABAC arrays:

```yaml
permissions:
  read:
    - rbac: PlatformAdmin # OR admin
    - rbac: AccountOwner # OR account owner
      abac:
        - AccountId #   AND matching account
    - abac:
        - ActorId # OR the actor themselves
```

#### Custom ABAC Field Mappings

Define how ABAC rules map to fields:

```yaml
abac:
  - name: AccountId
    field: accountId # Check if actor's accountId matches object's accountId
  - name: GroupId
    field: groupId
  - name: OwnerId
    field: ownerId
```

#### Special ABAC Rules

| Rule        | Description                                                        |
| ----------- | ------------------------------------------------------------------ |
| `ActorId`   | Checks if the actor's ID matches the object's `id` field           |
| `AccountId` | Checks if the actor's `accountId` matches the object's `accountId` |
| `GroupId`   | Checks if the actor's `groupId` matches the object's `groupId`     |

#### Field-Level Permissions

Apply permissions to individual fields:

```yaml
fields:
  - name: billingInfo
    type: BillingInfo
    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
      write:
        - rbac: PlatformAdmin # Only admin can modify billing info
```

#### Global Permission Definitions

Define roles and permissions in `permissions.yaml`:

```yaml
permissions:
  rbac:
    - name: PlatformAdmin
      customPermissions:
        - Platform
        - PlatformWrite

    - name: AccountOwner
      extends:
        - AccountAdmin # Inherits AccountAdmin permissions
      abac:
        - AccountId
      customPermissions:
        - Account

    - name: AccountAdmin
      abac:
        - AccountId
      customPermissions:
        - Account

  abac:
    - name: AccountId
      fieldType: Ref<Account>
    - name: GroupId
      fieldType: Ref<Group>

  roleGroups:
    - name: AdminDashboardLevel
      roles:
        - PlatformAdmin
        - AccountOwner
        - AccountAdmin
```

**Generated permission code:**

- Permission checking in `{model}/permissions.go`
- HTTP middleware for automatic enforcement
- Client-side permission utilities in TypeScript
- React hooks for conditional rendering

### Field Validation

Add validation rules that are enforced on both backend and frontend.

#### Available Validation Rules

| Rule       | Type    | Description                  | Example           |
| ---------- | ------- | ---------------------------- | ----------------- |
| `required` | boolean | Field must be present        | `required: true`  |
| `email`    | boolean | Must be valid email          | `email: true`     |
| `min`      | number  | Minimum value/length         | `min: 5`          |
| `max`      | number  | Maximum value/length         | `max: 100`        |
| `minItems` | number  | Minimum array length         | `minItems: 1`     |
| `maxItems` | number  | Maximum array length         | `maxItems: 10`    |
| `regex`    | string  | Pattern matching             | `regex: "^[A-Z]"` |
| `unique`   | boolean | Must be unique in collection | `unique: true`    |

#### Examples

**String Validation:**

```yaml
fields:
  - name: email
    type: string
    validate:
      required: true
      email: true

  - name: username
    type: string
    validate:
      required: true
      min: 3
      max: 20
      regex: "^[a-zA-Z0-9_]+$"
```

**Number Validation:**

```yaml
fields:
  - name: age
    type: int
    validate:
      required: true
      min: 0
      max: 120

  - name: price
    type: float64
    validate:
      required: true
      min: 0.01
```

**Array Validation:**

```yaml
fields:
  - name: tags
    type: List<string>
    validate:
      minItems: 1
      maxItems: 10

  - name: roles
    type: List<ActorRole>
    validate:
      required: true
      minItems: 1
```

**Unique Field:**

```yaml
fields:
  - name: email
    type: string
    validate:
      required: true
      email: true
      unique: true # Enforced at database level
```

**Where validation is enforced:**

- Go backend: Before database operations
- TypeScript: Form validation
- MongoDB: Unique constraints via indexes

### Actors and Authentication

Mark an object as an "actor" to enable it for authentication and authorization.

#### Actor Configuration

```yaml
actor:
  id: id # Field to use as actor ID
  username: "{email}" # Login username (supports templates)
  name: "{firstName} {lastName}" # Display name (supports templates)
  adminName: "{email} ({id})" # Admin interface display
  language: "{language}" # User's language preference
```

#### Template Syntax

Use `{fieldName}` to interpolate field values:

```yaml
actor:
  username: "{email}" # Uses email field
  name: "{firstName} {lastName}" # Combines multiple fields
  adminName: "{accountId} - {email}" # Custom format for admins
```

#### Example: User Actor

```yaml
objects:
  - name: User
    fields:
      - name: email
        type: string
      - name: firstName
        type: string
      - name: lastName
        type: string
      - name: language
        type: string

    actor:
      id: id
      username: "{email}"
      name: "{firstName} {lastName}"
      language: "{language}"

    collection:
      - type: mongo
        name: users
```

**What gets generated:**

- Actor interface implementations
- Authentication utilities
- Permission context (current actor)
- Role checking functions

### Enums

Define enumeration types for fields with fixed sets of values.

#### String Enum

```yaml
enums:
  - name: Role
    type: string
    values:
      - PlatformAdmin
      - AccountOwner
      - AccountAdmin
      - User
      - Guest
```

#### Integer Enum

```yaml
enums:
  - name: Priority
    type: int
    values:
      - Low: 1
      - Medium: 2
      - High: 3
      - Critical: 4
```

#### Using Enums in Objects

```yaml
objects:
  - name: Task
    fields:
      - name: priority
        type: Priority # References Priority enum
        validate:
          required: true
      - name: assigneeRole
        type: Role # References Role enum
```

**Generated code:**

- Go: Type-safe enum with constants and validation
- TypeScript: String union types or number enums
- Kotlin: Sealed classes or enums

**Go Example:**

```go
package enum_role

type Value string

const (
    PlatformAdmin Value = "PlatformAdmin"
    AccountOwner  Value = "AccountOwner"
    AccountAdmin  Value = "AccountAdmin"
    User          Value = "User"
    Guest         Value = "Guest"
)
```

**TypeScript Example:**

```typescript
export type Role =
  | "PlatformAdmin"
  | "AccountOwner"
  | "AccountAdmin"
  | "User"
  | "Guest";
```

### Errors and Events

#### Custom Errors

Define application-specific errors in `errors.yaml`:

```yaml
errors:
  - name: InvalidCredentials
    code: AUTH_001
    message: "Invalid username or password"

  - name: InsufficientPermissions
    code: AUTH_002
    message: "You do not have permission to perform this action"

  - name: ResourceNotFound
    code: DATA_001
    message: "The requested resource was not found"
```

**Generated:**

- Error types with codes
- Automatic HTTP status mapping
- Client-side error handling utilities

#### Events

Define domain events in `events.yaml`:

```yaml
eventPermissions:
  read:
    - rbac: PlatformAdmin
    - rbac: PlatformObserver
  write:
    - rbac: PlatformDeveloper

events:
  - UserCreated
  - UserUpdated
  - UserDeleted
  - AccountActivated
  - AccountDeactivated
  - PaymentProcessed
  - PaymentFailed
```

**Generated:**

- Event type enums
- Event logging utilities
- Event permission checking
- Event storage models

## Generated Code Reference

### Go Backend Structure

The generator creates a comprehensive Go package structure:

```
backend/models/
├── api/                              # Common API interfaces
│   ├── model.go                      # Base model interface
│   └── mongo.go                      # MongoDB interface
│
├── enum_role/                        # Generated enums
│   └── value.go                      # Enum constants and utilities
│
├── permissions/                      # Permission system
│   ├── actor.go                      # Actor interfaces
│   └── super.go                      # Super admin utilities
│
├── permissions_api/                  # Permission API
│   └── fetch_actor.go                # Actor fetching utilities
│
├── http_server/                      # HTTP routing
│   └── routes.go                     # Route registration
│
├── coded_error/                      # Error handling
│   └── error.go                      # Custom error types
│
├── event/                            # Event system
│   └── interface.go                  # Event interfaces
│
├── utils/                            # Utility functions
│   ├── conv.go                       # Conversion utilities
│   └── search.go                     # Search utilities
│
└── user/                            # Example: User model
    ├── model.go                      # Core model struct and methods
    ├── projection.go                 # Field projection utilities
    ├── mongo.go                      # MongoDB integration
    ├── http.go                       # HTTP handler registration
    ├── permissions.go                # Permission implementations
    │
    ├── user_api/                     # API layer
    │   ├── model.go                  # API interface
    │   ├── with_permissions.go       # Permission-wrapped API
    │   ├── unimplemented.go          # Unimplemented stubs
    │   ├── http.go                   # HTTP API implementation
    │   └── mongo.go                  # MongoDB API implementation
    │
    ├── user_mongo/                   # Database layer
    │   ├── collection.go             # Collection and index management
    │   ├── model.go                  # MongoDB record types
    │   ├── lookup.go                 # Lookup operations (with joins)
    │   ├── search.go                 # Search/query operations
    │   ├── save.go                   # Create/update operations
    │   └── delete.go                 # Delete operations
    │
    └── user_http/                    # HTTP layer
        ├── handlers.go               # HTTP request handlers
        └── routes.go                 # Route definitions
```

#### Key Generated Files Explained

**`model.go`** - Core model definition:

```go
package user

type Model struct {
    Id        string
    AccountId string
    Email     string
    FirstName string
    LastName  string
    Created   actor_trace.Model
    Updated   actor_trace.Model
}

// Conversion to/from MongoDB records
func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error)
func (m *Model) FromMongoRecord(r MongoRecord) error
```

**`projection.go`** - Field selection:

```go
type Projection struct {
    Id        bool
    AccountId bool
    Email     bool
    FirstName bool
    LastName  bool
    // ... etc
}

func AllFields() Projection {
    return Projection{
        Id: true,
        AccountId: true,
        // ... all fields true
    }
}
```

**`user_mongo/collection.go`** - Index management:

```go
package user_mongo

const CollectionName = "users"

func CreateIndexes(ctx context.Context, db *mongo.Database) error {
    // Creates all defined indexes
}
```

**`user_mongo/search.go`** - Query operations:

```go
func Search(ctx context.Context, db *mongo.Database, params SearchParams) ([]user.Model, error)
func Count(ctx context.Context, db *mongo.Database, query user_api.UserSearchQuery) (int64, error)
```

**`user_http/handlers.go`** - HTTP handlers:

```go
func HandleGet(api user_api.Model) http.HandlerFunc
func HandlePost(api user_api.Model) http.HandlerFunc
func HandlePatch(api user_api.Model) http.HandlerFunc
func HandleDelete(api user_api.Model) http.HandlerFunc
```

**`permissions.go`** - Permission checks:

```go
func CanRead(actor permissions.Actor, obj *Model) bool
func CanWrite(actor permissions.Actor, obj *Model) bool
```

### TypeScript Frontend Structure

Generated TypeScript files provide complete type safety and API integration:

```
frontend/lib/
├── api/                              # API utilities
│   ├── model.ts                      # Base API client
│   └── errors.ts                     # Error types
│
├── model/                            # TypeScript models
│   ├── user-model.ts                 # User type definitions
│   ├── user-api.ts                   # User API types (queries, filters)
│   ├── role-enum.ts                  # Enum definitions
│   └── ...
│
├── endpoints/                        # API endpoints
│   ├── user-endpoints.ts             # User CRUD functions
│   └── ...
│
├── permissions/                      # Permission utilities
│   ├── actor.ts                      # Actor type
│   ├── user-can-access.ts            # User permission checking
│   └── ...
│
└── react/                           # React integration
    ├── api.tsx                       # API context provider
    │
    ├── tanstack-query/               # React Query hooks
    │   ├── user-queries.ts           # User queries and mutations
    │   └── ...
    │
    ├── form-state/                   # Form management
    │   ├── common.ts                 # Common form utilities
    │   ├── user-form-state.ts        # User form state
    │   └── ...
    │
    └── mui/                          # Material-UI integration
        ├── user-data-grid.ts         # DataGrid configuration
        ├── user-search-selector.ts   # Autocomplete/selector
        └── ...
```

#### Key Generated Files Explained

**`model/user-model.ts`** - Type definitions:

```typescript
export type User = {
  id?: string;
  accountId?: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  created?: ActorTrace;
  updated?: ActorTrace;
};

export type UserProjection = {
  id?: boolean;
  accountId?: boolean;
  email?: boolean;
  // ... field flags
};

export type UserSortParams = {
  email?: -1 | 1;
  createdAt?: -1 | 1;
  // ... sortable fields
};
```

**`model/user-api.ts`** - Query and filter types:

```typescript
export type UserSearchQuery = {
  // ID filters
  idEq?: string;
  idIn?: string[];
  idNin?: string[];

  // String field filters
  emailEq?: string;
  emailLike?: string;
  emailExists?: boolean;

  // Ref filters
  accountIdEq?: string;
  accountIdIn?: string[];

  // ... all searchable fields
};

export type UserWithRefs = {
  user: User;
  account?: Account;
  group?: Group;
  // ... related objects
};
```

**`endpoints/user-endpoints.ts`** - API functions:

```typescript
export async function getUser(
  api: ApiContext,
  query: SelectUserByIdQuery,
  projection?: UserProjection
): Promise<User> {
  // GET /users/:id
}

export async function searchUsers(
  api: ApiContext,
  query: UserSearchQuery,
  projection?: UserProjection,
  sort?: UserSortParams,
  limit?: number,
  skip?: number
): Promise<User[]> {
  // GET /users with query params
}

export async function createUser(api: ApiContext, user: User): Promise<User> {
  // POST /users
}

export async function updateUser(
  api: ApiContext,
  id: string,
  updates: Partial<User>
): Promise<User> {
  // PATCH /users/:id
}

export async function deleteUser(api: ApiContext, id: string): Promise<void> {
  // DELETE /users/:id
}
```

**`react/tanstack-query/user-queries.ts`** - React Query hooks:

```typescript
export function useUserQueries() {
  const api = useApi();

  return {
    useGetUser: (query: SelectUserByIdQuery) =>
      useQuery({
        queryKey: ["user", query.id],
        queryFn: () => getUser(api, query),
      }),

    useSearchUsers: (query: UserSearchQuery) =>
      useQuery({
        queryKey: ["users", "search", query],
        queryFn: () => searchUsers(api, query),
      }),

    useCreateUser: () =>
      useMutation({
        mutationFn: (user: User) => createUser(api, user),
        onSuccess: () => {
          queryClient.invalidateQueries({ queryKey: ["users"] });
        },
      }),

    useUpdateUser: () =>
      useMutation({
        mutationFn: ({ id, updates }: { id: string; updates: Partial<User> }) =>
          updateUser(api, id, updates),
        onSuccess: () => {
          queryClient.invalidateQueries({ queryKey: ["users"] });
        },
      }),

    useDeleteUser: () =>
      useMutation({
        mutationFn: (id: string) => deleteUser(api, id),
        onSuccess: () => {
          queryClient.invalidateQueries({ queryKey: ["users"] });
        },
      }),
  };
}
```

**`react/form-state/user-form-state.ts`** - Form state management:

```typescript
export function useUserFormState(initialUser?: User) {
  const [user, setUser] = useState<User>(initialUser || {});
  const [errors, setErrors] = useState<Record<string, string>>({});

  const validate = () => {
    const newErrors: Record<string, string> = {};

    // Generated validation logic
    if (!user.email) {
      newErrors.email = "Email is required";
    } else if (!isValidEmail(user.email)) {
      newErrors.email = "Invalid email format";
    }

    // ... more validation

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  return { user, setUser, errors, validate };
}
```

**`permissions/user-can-access.ts`** - Permission checking:

```typescript
export function canReadUser(actor: Actor, user: User): boolean {
  // Platform admin can read all
  if (actor.roles.includes("PlatformAdmin")) {
    return true;
  }

  // Account owner can read their account's users
  if (
    actor.roles.includes("AccountOwner") &&
    actor.accountId === user.accountId
  ) {
    return true;
  }

  // User can read themselves
  if (actor.id === user.id) {
    return true;
  }

  return false;
}

export function canWriteUser(actor: Actor, user: User): boolean {
  // Similar permission logic for write
}
```

**`mui/user-data-grid.ts`** - Material-UI DataGrid config:

```typescript
export const userDataGridColumns: GridColDef[] = [
  {
    field: "id",
    headerName: "ID",
    width: 200,
  },
  {
    field: "email",
    headerName: "Email",
    width: 200,
  },
  {
    field: "firstName",
    headerName: "First Name",
    width: 150,
  },
  {
    field: "lastName",
    headerName: "Last Name",
    width: 150,
  },
  // ... all fields with appropriate renderers
];
```

### Kotlin Mobile Structure

Generated Kotlin files for mobile applications:

```
mobile/app/src/main/java/com/example/models/
├── User.kt                           # User data class
├── UserApi.kt                        # API client
├── UserPermissions.kt                # Permission checking
├── Role.kt                          # Enum definitions
└── ...
```

**Example Kotlin model:**

```kotlin
@Serializable
data class User(
    val id: String? = null,
    val accountId: String? = null,
    val email: String? = null,
    val firstName: String? = null,
    val lastName: String? = null,
    val created: ActorTrace? = null,
    val updated: ActorTrace? = null
)

@Serializable
data class UserSearchQuery(
    val idEq: String? = null,
    val idIn: List<String>? = null,
    val emailEq: String? = null,
    val emailLike: String? = null,
    // ... filters
)
```

## Advanced Features

### Object References

Create relationships between objects using reference types.

#### Simple Reference (Foreign Key)

```yaml
objects:
  - name: Post
    fields:
      - name: authorId
        type: Ref<User> # Foreign key to User
      - name: title
        type: string
      - name: content
        type: string
```

**Generated behavior:**

- Stores ObjectID in MongoDB
- Automatic validation (user must exist)
- Lookup queries that join with referenced object

**Usage in Go:**

```go
post := post.Model{
    AuthorId: "507f1f77bcf86cd799439011",  // User's ID
    Title:    "My Post",
    Content:  "Post content...",
}
```

**Lookup with references:**

```go
// Automatically joins with User
postsWithAuthors := api.LookupPosts(ctx, query)
// postsWithAuthors[0].Author contains the full User object
```

#### List of References

```yaml
fields:
  - name: memberIds
    type: List<Ref<User>> # Array of user IDs
```

**Generated:**

- Array of ObjectIDs in MongoDB
- Lookup queries that fetch all referenced users

#### Reference with Related Objects

The generator automatically creates bidirectional relationships:

```yaml
# Post has an author
objects:
  - name: Post
    fields:
      - name: authorId
        type: Ref<User>
# Lookup includes related posts
# GET /users/:id/lookup returns user with their posts
```

### Parent-Child Relationships

Create hierarchical structures with `ParentRef`.

#### Example: Account → Group → User Hierarchy

```yaml
objects:
  - name: Account
    fields:
      - name: name
        type: string
    collection:
      - type: mongo
        name: accounts

  - name: Group
    fields:
      - name: accountId
        type: ParentRef<Account> # Group belongs to Account
      - name: name
        type: string
    collection:
      - type: mongo
        name: groups

  - name: User
    fields:
      - name: accountId
        type: ParentRef<Account> # User belongs to Account
      - name: groupId
        type: ParentRef<Group> # User belongs to Group
      - name: email
        type: string
    collection:
      - type: mongo
        name: users
```

**Generated behavior:**

- Enforces referential integrity
- Automatic cascade queries (get all children)
- Permission inheritance patterns

**Usage:**

```typescript
// Fetch account with all groups and users
const accountWithChildren = await lookupAccount(api, { id: accountId });
// Result includes:
// - account: Account
// - groups: Group[]
// - users: User[] (from all groups)
```

#### ParentRef Benefits

1. **Enforced Hierarchy**: Cannot create orphaned records
2. **Cascade Queries**: Fetch entire subtrees efficiently
3. **Permission Propagation**: Parent permissions can apply to children
4. **Automatic Indexes**: Optimized for hierarchical queries

### Custom ABAC Rules

Define custom attribute-based access control logic.

#### Basic ABAC Definition

```yaml
objects:
  - name: Document
    fields:
      - name: ownerId
        type: Ref<User>
      - name: departmentId
        type: Ref<Department>
      - name: isPublic
        type: bool

    # Define custom ABAC rules
    abac:
      - name: IsOwner
        field: ownerId # Check if actor.id == document.ownerId

      - name: InDepartment
        field: departmentId # Check if actor.departmentId == document.departmentId

    permissions:
      read:
        - rbac: Admin # Admins read all
        - abac:
            - IsOwner # Owner can read
        - abac:
            - InDepartment # Department members can read

      write:
        - rbac: Admin
        - abac:
            - IsOwner # Only owner can write
```

#### Complex Multi-Field ABAC

```yaml
abac:
  - name: SameAccountAndGroup
    field: accountId # Must match accountId
    additionalFields:
      - groupId # AND must match groupId

permissions:
  write:
    - rbac: GroupAdmin
      abac:
        - SameAccountAndGroup # Admin only for their specific group
```

#### ABAC with List Fields

```yaml
fields:
  - name: editorIds
    type: List<Ref<User>>

abac:
  - name: IsEditor
    field: editorIds # Check if actor.id is in editorIds array

permissions:
  write:
    - abac:
        - IsEditor # Can write if in editors list
```

**Generated permission check:**

```go
func CanWrite(actor permissions.Actor, doc *Model) bool {
    // Check if actor is in editorIds
    for _, editorId := range doc.EditorIds {
        if editorId == actor.GetId() {
            return true
        }
    }
    return false
}
```

### Field-Level Permissions

Apply different permissions to specific fields.

#### Example: Sensitive Financial Data

```yaml
objects:
  - name: User
    fields:
      - name: email
        type: string
        # Email visible to account admins

      - name: billingInfo
        type: BillingInfo
        permissions:
          read:
            - rbac: PlatformAdmin
            - rbac: AccountOwner # Only owner sees billing
          write:
            - rbac: PlatformAdmin # Only admin can modify

      - name: internalNotes
        type: string
        permissions:
          read:
            - rbac: PlatformAdmin # Only internal staff
          write:
            - rbac: PlatformAdmin
```

#### Nested Object with Field Permissions

```yaml
# obj_billing_info.yaml
objects:
  - name: BillingInfo
    fields:
      - name: email
        type: string
      - name: phone
        type: string
      - name: creditCardLast4
        type: string
        permissions:
          read:
            - rbac: PlatformAdmin
            - rbac: AccountOwner
      - name: country
        type: string
        permissions:
          write:
            - rbac: PlatformAdmin # Country locked after creation
```

**Generated behavior:**

- Automatic field filtering in API responses
- Projection-aware queries (exclude forbidden fields)
- Validation prevents unauthorized field updates

**API behavior:**

```typescript
// Non-owner gets user without billingInfo
const user = await getUser(api, { id: userId });
// user.billingInfo === undefined

// Owner gets full user including billingInfo
const userForOwner = await getUser(apiAsOwner, { id: userId });
// userForOwner.billingInfo === { email: "...", phone: "..." }
```

### Immutable Fields

Mark fields that cannot be changed after creation.

```yaml
fields:
  - name: accountId
    type: ParentRef<Account>
    immutable: true # Cannot be changed after creation

  - name: createdAt
    type: timestamp
    immutable: true

  - name: userId
    type: Ref<User>
    immutable: true # User assignment is permanent
```

**Generated behavior:**

- Go: Validation prevents updates to immutable fields
- TypeScript: Type hints indicate immutability
- HTTP: PATCH requests ignore immutable fields

**Update validation:**

```go
func Update(ctx context.Context, id string, updates Model) error {
    existing, err := Get(ctx, id)
    if err != nil {
        return err
    }

    // Validation: reject if immutable fields changed
    if updates.AccountId != "" && updates.AccountId != existing.AccountId {
        return errors.New("accountId is immutable")
    }

    // ... proceed with update
}
```

## Practical Examples

### Example 1: E-commerce Product Catalog

```yaml
# model_product.yaml
enums:
  - name: ProductStatus
    type: string
    values:
      - Draft
      - Published
      - Archived

objects:
  - name: Product
    fields:
      - name: accountId
        type: ParentRef<Account>
        immutable: true
      - name: name
        type: string
        validate:
          required: true
          min: 3
          max: 200
      - name: sku
        type: string
        validate:
          required: true
          unique: true
          regex: "^[A-Z0-9-]+$"
      - name: description
        type: string
        validate:
          max: 5000
      - name: price
        type: float64
        validate:
          required: true
          min: 0
      - name: status
        type: ProductStatus
        validate:
          required: true
      - name: categoryIds
        type: List<Ref<Category>>
      - name: tags
        type: List<string>
        validate:
          maxItems: 20
      - name: imageUrls
        type: List<string>
      - name: inStock
        type: bool
      - name: stockCount
        type: int
        validate:
          min: 0

    collection:
      - type: mongo
        name: products

    indexes:
      - name: sku_unique
        unique: true
        fields:
          - name: sku
      - name: account_status
        fields:
          - name: accountId
          - name: status
      - name: search_index
        fields:
          - name: name
          - name: tags

    http:
      endpoint: products
      methods:
        - GET
        - POST
        - PATCH
        - DELETE

    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId
        - rbac: AccountAdmin
          abac:
            - AccountId
      write:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId

    abac:
      - name: AccountId
        field: accountId
```

**Usage:**

```typescript
// Search for published products
const products = await searchProducts(api, {
  statusEq: "Published",
  inStockEq: true,
  priceGte: 10,
  priceLte: 100,
  tagsIn: ["electronics", "gadgets"],
});

// Create product
const newProduct = await createProduct(api, {
  accountId: currentAccount.id,
  name: "Smart Watch",
  sku: "WATCH-001",
  price: 299.99,
  status: "Draft",
  inStock: true,
  stockCount: 50,
});
```

### Example 2: Project Management System

```yaml
# model_project.yaml
enums:
  - name: ProjectStatus
    type: string
    values:
      - Planning
      - Active
      - OnHold
      - Completed
      - Cancelled

  - name: TaskPriority
    type: int
    values:
      - Low: 1
      - Medium: 2
      - High: 3
      - Critical: 4

objects:
  - name: Project
    fields:
      - name: accountId
        type: ParentRef<Account>
        immutable: true
      - name: name
        type: string
        validate:
          required: true
      - name: description
        type: string
      - name: status
        type: ProjectStatus
        validate:
          required: true
      - name: ownerId
        type: Ref<User>
        validate:
          required: true
      - name: memberIds
        type: List<Ref<User>>
      - name: startDate
        type: timestamp
      - name: endDate
        type: timestamp
      - name: budget
        type: float64
        permissions:
          read:
            - rbac: PlatformAdmin
            - rbac: AccountOwner
            - abac:
                - IsOwner
          write:
            - rbac: PlatformAdmin
            - rbac: AccountOwner
      - name: created
        type: ActorTrace
      - name: updated
        type: ActorTrace

    collection:
      - type: mongo
        name: projects

    indexes:
      - name: account_status
        fields:
          - name: accountId
          - name: status
      - name: owner
        fields:
          - name: ownerId
      - name: members
        fields:
          - name: memberIds

    http:
      endpoint: projects
      methods:
        - GET
        - POST
        - PATCH
        - DELETE

    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId
        - abac:
            - IsOwner
        - abac:
            - IsMember
      write:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId
        - abac:
            - IsOwner

    abac:
      - name: AccountId
        field: accountId
      - name: IsOwner
        field: ownerId
      - name: IsMember
        field: memberIds

  - name: Task
    fields:
      - name: projectId
        type: ParentRef<Project>
        immutable: true
      - name: title
        type: string
        validate:
          required: true
      - name: description
        type: string
      - name: assigneeId
        type: Ref<User>
      - name: priority
        type: TaskPriority
        validate:
          required: true
      - name: status
        type: string
        validate:
          required: true
      - name: dueDate
        type: timestamp
      - name: completed
        type: bool
      - name: created
        type: ActorTrace
      - name: updated
        type: ActorTrace

    collection:
      - type: mongo
        name: tasks

    indexes:
      - name: project_status
        fields:
          - name: projectId
          - name: status
      - name: assignee
        fields:
          - name: assigneeId
          - name: completed

    http:
      endpoint: tasks
      methods:
        - GET
        - POST
        - PATCH
        - DELETE

    permissions:
      read:
        - rbac: PlatformAdmin
        - abac:
            - InProject # Can read if part of parent project
      write:
        - rbac: PlatformAdmin
        - abac:
            - InProject

    abac:
      - name: InProject
        field: projectId # Check via parent project permissions
```

**Usage:**

```typescript
// React component with generated hooks
function ProjectBoard() {
  const { useSearchProjects, useCreateProject } = useProjectQueries();
  const { useSearchTasks, useUpdateTask } = useTaskQueries();

  // Fetch active projects
  const { data: projects } = useSearchProjects({
    statusEq: "Active",
    memberIdsIn: [currentUser.id],
  });

  // Fetch tasks for a project
  const { data: tasks } = useSearchTasks(
    {
      projectIdEq: selectedProject?.id,
      completedEq: false,
    },
    {
      sort: { priority: -1, dueDate: 1 },
    }
  );

  // Update task
  const { mutate: updateTask } = useUpdateTask();

  const handleCompleteTask = (taskId: string) => {
    updateTask({ id: taskId, updates: { completed: true } });
  };

  return <div>{/* Render projects and tasks */}</div>;
}
```

### Example 3: Multi-Tenant SaaS Platform

```yaml
# permissions.yaml
permissions:
  rbac:
    - name: PlatformAdmin
      customPermissions:
        - Platform
        - PlatformWrite

    - name: AccountOwner
      extends:
        - AccountAdmin
      abac:
        - AccountId
      customPermissions:
        - Account
        - Billing

    - name: AccountAdmin
      abac:
        - AccountId
      customPermissions:
        - Account

    - name: TeamMember
      abac:
        - AccountId
        - TeamId

# model_account.yaml
objects:
  - name: Account
    fields:
      - name: name
        type: string
        validate:
          required: true
      - name: slug
        type: string
        validate:
          required: true
          unique: true
          regex: "^[a-z0-9-]+$"
      - name: plan
        type: string
        validate:
          required: true
      - name: isActive
        type: bool
      - name: settings
        type: Map<string,string>
      - name: created
        type: ActorTrace

    collection:
      - type: mongo
        name: accounts

    indexes:
      - name: slug_unique
        unique: true
        fields:
          - name: slug

    actor:
      id: id
      name: "{name}"

    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId
        - rbac: AccountAdmin
          abac:
            - AccountId
      write:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId

    abac:
      - name: AccountId
        field: id

# model_team.yaml
objects:
  - name: Team
    fields:
      - name: accountId
        type: ParentRef<Account>
        immutable: true
      - name: name
        type: string
        validate:
          required: true
      - name: leaderIds
        type: List<Ref<User>>
      - name: memberIds
        type: List<Ref<User>>

    collection:
      - type: mongo
        name: teams

    permissions:
      read:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId
        - rbac: AccountAdmin
          abac:
            - AccountId
        - abac:
            - IsMember
      write:
        - rbac: PlatformAdmin
        - rbac: AccountOwner
          abac:
            - AccountId

    abac:
      - name: AccountId
        field: accountId
      - name: IsMember
        field: memberIds
```

## Development Guide

### Adding New Field Types

To support a new field type across the entire stack:

1. **Add type to parser** (`platform/types/field.go`):

```go
const (
    FieldTypeString    = "string"
    FieldTypeInt       = "int"
    // Add your new type
    FieldTypeMyNewType = "mynewtype"
)
```

2. **Add Go type mapping** (`platform/templates/go/utils.go`):

```go
func GoType(field types.Field) string {
    switch field.Type {
    case types.FieldTypeMyNewType:
        return "MyNewGoType"
    // ...
    }
}
```

3. **Add TypeScript type mapping** (`platform/templates/ts/utils.go`):

```go
func TSType(field types.Field) string {
    switch field.Type {
    case types.FieldTypeMyNewType:
        return "MyNewTSType"
    // ...
    }
}
```

4. **Update MongoDB conversion** (in model templates):

- Add conversion logic in `ToMongoRecord`
- Add parsing logic in `FromMongoRecord`

5. **Update validation** (`platform/templates/go/utils_template_model.go`):

```go
// Add validation rules for your new type
```

### Creating Custom Templates

#### 1. Create Template File

Create a new file in `platform/templates/go/` or `platform/templates/ts/`:

```go
// platform/templates/go/template_my_custom.go
package model_template_go

const MyCustomTemplate = `
// Custom generated code for {{ .Object.Name }}
package {{ .PackageName }}

func CustomFunction() {
    // Your template logic
}
`
```

#### 2. Register Template

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

#### 3. Use Template Functions

Available template functions (defined in `template_funcs.go`):

```go
{{ .Object.Name | ToLower }}           // Convert to lowercase
{{ .Object.Name | ToCamel }}           // Convert to camelCase
{{ .Object.Name | ToSnake }}           // Convert to snake_case
{{ .Field.Type | GoType }}             // Get Go type
{{ .Field.Type | TSType }}             // Get TypeScript type
{{ if .Object.HasCollection }}...{{ end }}  // Conditional
```

### Extending the Permission System

#### Add New Permission Type

1. **Define in `permissions.yaml`**:

```yaml
permissions:
  rbac:
    - name: MyNewRole
      customPermissions:
        - MyCustomPermission
```

2. **Update permission templates**:

- `template_permissions_actor.go` - Actor interface
- `template_obj__model_permissions.go` - Permission checking
- `template_permissions_obj__can_access.go` - TypeScript utilities

3. **Implement custom logic**:

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

### Running Tests

```bash
# Test the generator
cd platform
go test ./...

# Test generated Go code
cd ../backend
go test ./models/...

# Test generated TypeScript
cd ../frontend
npm test
```

### Debugging Generated Code

1. **Enable verbose output**:

```bash
go run main.go --specDir=../platform_models --verbose
```

2. **Check validation errors**:
   The generator validates specs before generation. Common errors:

- Unknown field types
- Circular references
- Invalid permission configurations
- Duplicate object names

3. **Inspect generated files**:

- Go files include package imports and type info
- TypeScript files are formatted with Prettier
- All files have generation timestamp comments

## Architecture & How It Works

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
2. **Review examples** - Look at existing model definitions in `platform_models/`
3. **Inspect generated code** - Generated files include comments and structure
4. **Check types** - Ensure field types match available types

---

**Generated Code Notice**: All generated files include a header comment:

```
// This file is auto-generated. DO NOT EDIT.
```

Never manually edit generated files. Instead, modify the YAML specifications and regenerate.
