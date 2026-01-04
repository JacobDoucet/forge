#!/bin/bash
# Build script for the example models
# This script generates Go code from the example YAML model specifications.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
FORGE_DIR="$(dirname "$SCRIPT_DIR")"

echo "Building forge..."
cd "$FORGE_DIR"
go build -o forge .

echo "Generating example models..."
./forge build \
    --specDir "$SCRIPT_DIR/models" \
    --goOutDir "$SCRIPT_DIR/generated" \
    --goPkgRoot "d3tech.com/platform/example/generated"

echo "Verifying generated code compiles..."
cd "$SCRIPT_DIR"
go build ./...

echo "✓ Example build completed successfully!"
