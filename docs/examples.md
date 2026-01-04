# Practical Examples

[← Back to README](../README.md)

This guide provides complete, real-world examples of Forge model definitions.

## Table of Contents

- [Example 1: E-commerce Product Catalog](#example-1-e-commerce-product-catalog)
- [Example 2: Project Management System](#example-2-project-management-system)
- [Example 3: Multi-Tenant SaaS Platform](#example-3-multi-tenant-saas-platform)

## Example 1: E-commerce Product Catalog

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

## Example 2: Project Management System

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

## Example 3: Multi-Tenant SaaS Platform

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

**Multi-tenant architecture features:**

- Account-scoped data isolation
- Role-based access with ABAC refinement
- Team hierarchies within accounts
- Flexible permission inheritance
