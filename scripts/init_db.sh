#!/bin/bash

# Navigate to project root
cd "$(dirname "$0")/.."

# Run the initialization script
echo "Initializing database with default data..."
go run cmd/migrate/init_data.go

echo "Done!"