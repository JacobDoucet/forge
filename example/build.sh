#!/bin/bash
# Build script for the example models
# This script generates code from the example YAML model specifications.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
FORGE_DIR="$(dirname "$SCRIPT_DIR")"

echo "Building forge..."
cd "$FORGE_DIR"
go build -o forge .

echo "Generating example models..."
cd "$SCRIPT_DIR"
"$FORGE_DIR/forge" build

echo "Verifying generated Go code compiles..."
cd "$SCRIPT_DIR/apps/backend"
go mod tidy
go build ./generated/...

echo "✓ Example build completed successfully!"
echo ""
echo "To run the full stack with Docker:"
echo "  cd $SCRIPT_DIR && docker-compose up --build"
echo ""
echo "Or run individually:"
echo "  Backend: cd apps/backend && go run ."
echo "  Frontend: cd apps/frontend && npm install && npm run dev"
