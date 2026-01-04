# Advanced Features

[← Back to README](../README.md)

This guide covers advanced features for complex data models and access control patterns.

## Table of Contents

- [Object References](#object-references)
- [Parent-Child Relationships](#parent-child-relationships)
- [Custom ABAC Rules](#custom-abac-rules)
- [Field-Level Permissions](#field-level-permissions)
- [Immutable Fields](#immutable-fields)

## Object References

Create relationships between objects using reference types.

### Simple Reference (Foreign Key)

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

### List of References

```yaml
fields:
  - name: memberIds
    type: List<Ref<User>> # Array of user IDs
```

**Generated:**

- Array of ObjectIDs in MongoDB
- Lookup queries that fetch all referenced users

### Reference with Related Objects

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

## Parent-Child Relationships

Create hierarchical structures with `ParentRef`.

### Example: Account → Group → User Hierarchy

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

### ParentRef Benefits

1. **Enforced Hierarchy**: Cannot create orphaned records
2. **Cascade Queries**: Fetch entire subtrees efficiently
3. **Permission Propagation**: Parent permissions can apply to children
4. **Automatic Indexes**: Optimized for hierarchical queries

## Custom ABAC Rules

Define custom attribute-based access control logic.

### Basic ABAC Definition

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

### Complex Multi-Field ABAC

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

### ABAC with List Fields

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

## Field-Level Permissions

Apply different permissions to specific fields.

### Example: Sensitive Financial Data

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

### Nested Object with Field Permissions

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

## Immutable Fields

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
