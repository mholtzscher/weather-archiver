DB_DRIVER := postgres
DB_STRING := "host=localhost user=postgres password=postgres dbname=formula-data sslmode=disable"

# Define the directory where migration files are stored
MIGRATIONS_DIR := sql/migrations

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migration-create name=<migration_name>"
	@echo "  make migration-up"
	@echo "  make migration-down"
	@echo "  make migration-status"
	@echo "  make sqlc-gen"
	@echo "  make buf-gen"
	@echo "  make run"

# Create a new migration
.PHONY: migration-create
migration-create:
	@if [ -z "$(name)" ]; then echo "Please provide a migration name like this: make create-migration name=your_migration"; exit 1; fi
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Migrate the DB to the most recent version available
.PHONY: migration-up
migration-up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) up

# Roll back the version by 1
.PHONY: migration-down
migration-down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) down

# Display the status of all migrations
.PHONY: migration-status
migration-status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) status

# Run the server
.PHONY: run
run: 
	go run cmd/server/server.go

# Generate the code for the sqlc
.PHONY: sqlc-gen
sqlc-gen:
	sqlc generate

# Generate code from the proto files
.PHONY: buf-gen
buf-gen:
	buf generate

# Generate mocks for testing
.PHONY: mock-gen
mock-gen:
	mockery

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run all gen commands
.PHONY: gen
gen: buf-gen migration-up sqlc-gen mock-gen 

# Run all 
.PHONY: all
all: buf-gen migration-up sqlc-gen mock-gen run
