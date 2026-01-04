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

echo "Verifying generated code compiles..."
go build ./generated/go/...

echo "✓ Example build completed successfully!"
