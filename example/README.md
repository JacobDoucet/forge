# Example Module

This directory contains an example project demonstrating Forge code generation with real working applications.

## Structure

```
example/
├── models/                  # YAML model specifications
│   ├── enum_*.yml           # Enum definitions
│   ├── object_*.yml         # Object/model definitions
│   ├── errors.yml           # Error definitions
│   ├── events.yml           # Event definitions
│   └── permissions.yml      # Permission definitions
├── apps/
│   ├── backend/             # Go API server
│   │   ├── main.go          # Server entry point
│   │   ├── generated/       # Forge-generated Go code
│   │   └── Dockerfile
│   └── frontend/            # React application
│       ├── src/             # React source code
│       ├── src/generated/   # Forge-generated TypeScript code
│       └── Dockerfile
├── docker-compose.yml       # Docker orchestration
├── .forge.yml               # Forge configuration
└── build.sh                 # Build script
```

## Quick Start

### Using Docker Compose

The easiest way to run the full stack:

```bash
# Generate code first
./build.sh

# Start all services (MongoDB, Backend, Frontend)
docker-compose up --build
```

Then open http://localhost:3000 in your browser.

### Development Mode

#### Backend

```bash
cd apps/backend
go mod tidy
MONGO_URI=mongodb://localhost:27017 go run .
```

#### Frontend

```bash
cd apps/frontend
npm install
npm run dev
```

## Code Generation

Run the build script to generate code from the YAML model specifications:

```bash
./build.sh
```

Or use the Taskfile from the repo root:

```bash
task example:build
```

This will:

1. Build the forge CLI tool
2. Generate Go code to `apps/backend/generated/`
3. Generate TypeScript code to `apps/frontend/src/generated/`
4. Verify the Go generated code compiles

## Example Model

The example model (`models/`) demonstrates:

- **Enums**: `TaskStatus`, `TaskPriority`
- **Objects with Collections**: `Task`, `Project` (MongoDB-backed)
- **Nested Objects**: `TaskComment` (embedded type)
- **HTTP Endpoints**: REST API configuration
- **Permissions**: RBAC configuration
- **Indexes**: MongoDB index definitions
- **Errors**: Custom error types
- **Events**: Event type definitions

## API Endpoints

The backend exposes the following REST endpoints:

- `GET/POST /tasks/` - List/Create tasks
- `GET/PUT/DELETE /tasks/{id}` - Get/Update/Delete a task
- `GET/POST /projects/` - List/Create projects
- `GET/PUT/DELETE /projects/{id}` - Get/Update/Delete a project
- `GET/POST /events/` - List/Create events

## Cleaning Up

To remove generated files and Docker resources:

```bash
# Remove generated code
rm -rf apps/backend/generated/
rm -rf apps/frontend/src/generated/

# Stop and remove Docker containers
docker-compose down -v
```
