# YAML Specification Guide

[← Back to README](../README.md)

This guide covers the complete YAML specification format for defining your data models, permissions, and generated code behavior.

## Table of Contents

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

## Object Definition

Objects are the core building blocks. Each object definition creates a complete set of model, API, database, and UI code.

### Basic Structure

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

### Complete Example

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

## Field Types Reference

### Primitive Types

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

### Reference Types

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

### Collection Types

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

### Special Types

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

### Enum Types

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

### Embedded Object Types

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
        type: BillingInfo # Embedded object, not a reference
```

## Collections

Collections define how objects are persisted in the database.

**Supported database types:**

- `mongo` - MongoDB
- `postgres` - PostgreSQL _(coming soon)_
- `sqlite` - SQLite _(coming soon)_

```yaml
collection:
  - type: mongo
    name: collection_name # Collection/table name (plural, snake_case recommended)
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

## Indexes

Define database indexes to optimize query performance. Indexes are applied to your configured database backend.

### Simple Index

```yaml
indexes:
  - name: email_idx
    fields:
      - name: email
        order: 1 # 1 = ascending, -1 = descending
```

### Unique Index

```yaml
indexes:
  - name: email_unique
    unique: true
    fields:
      - name: email
```

### Compound Index

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

### Nested Field Index

```yaml
indexes:
  - name: created_at
    fields:
      - name: created.at # Index on nested field
        order: -1 # Descending (newest first)
```

### Multi-Field Index Example

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

- Index creation utilities for your database backend
- Automatic index creation on application startup
- Index management utilities

## HTTP Endpoints

Define which HTTP methods are available for your object.

### Basic Configuration

```yaml
http:
  endpoint: users # URL segment (defaults to lowercase object name)
  methods:
    - GET # GET /users/:id
    - POST # POST /users
    - PATCH # PATCH /users/:id
    - DELETE # DELETE /users/:id
```

### Generated Routes

| Method   | Route           | Description               |
| -------- | --------------- | ------------------------- |
| `GET`    | `/endpoint/:id` | Fetch single object by ID |
| `POST`   | `/endpoint`     | Create new object         |
| `PATCH`  | `/endpoint/:id` | Update existing object    |
| `DELETE` | `/endpoint/:id` | Delete object             |

**Additional auto-generated endpoints:**

- `GET /endpoint` - Search/list with query parameters
- `GET /endpoint/:id/lookup` - Fetch with related objects

### Custom Endpoint Names

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

## Permissions System

The generator supports sophisticated Role-Based Access Control (RBAC) and Attribute-Based Access Control (ABAC).

### Permission Structure

```yaml
permissions:
  read: [] # Who can read this object
  write: [] # Who can create/update/delete this object
```

### RBAC (Role-Based Access Control)

Simple role-based permissions:

```yaml
permissions:
  read:
    - rbac: PlatformAdmin # PlatformAdmin role can read
    - rbac: AccountOwner # AccountOwner role can read
  write:
    - rbac: PlatformAdmin # Only PlatformAdmin can write
```

### ABAC (Attribute-Based Access Control)

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

### Multiple ABAC Rules (AND Logic)

```yaml
permissions:
  read:
    - rbac: GroupAdmin
      abac:
        - AccountId # Must match account ID
        - GroupId # AND must match group ID
```

### Permission Rule Evaluation

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

### Custom ABAC Field Mappings

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

### Special ABAC Rules

| Rule        | Description                                                        |
| ----------- | ------------------------------------------------------------------ |
| `ActorId`   | Checks if the actor's ID matches the object's `id` field           |
| `AccountId` | Checks if the actor's `accountId` matches the object's `accountId` |
| `GroupId`   | Checks if the actor's `groupId` matches the object's `groupId`     |

### Field-Level Permissions

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

### Global Permission Definitions

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

## Field Validation

Add validation rules that are enforced on both backend and frontend.

### Available Validation Rules

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

### Examples

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

## Actors and Authentication

Mark an object as an "actor" to enable it for authentication and authorization.

### Actor Configuration

```yaml
actor:
  id: id # Field to use as actor ID
  username: "{email}" # Login username (supports templates)
  name: "{firstName} {lastName}" # Display name (supports templates)
  adminName: "{email} ({id})" # Admin interface display
  language: "{language}" # User's language preference
```

### Template Syntax

Use `{fieldName}` to interpolate field values:

```yaml
actor:
  username: "{email}" # Uses email field
  name: "{firstName} {lastName}" # Combines multiple fields
  adminName: "{accountId} - {email}" # Custom format for admins
```

### Example: User Actor

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

## Enums

Define enumeration types for fields with fixed sets of values.

### String Enum

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

### Integer Enum

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

### Using Enums in Objects

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

## Errors and Events

### Custom Errors

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

### Events

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
