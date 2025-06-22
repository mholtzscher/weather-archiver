# Variablejust
DB_DRIVER := "postgres"
DB_STRING := "host=localhost user=postgres password=postgres dbname=weather-archiver sslmode=disable"
MIGRATIONS_DIR := "sql/migrations"

help:
  @echo "Usage:"
  @echo "  just migration-create name=<migration_name>"
  @echo "  just migration-up"
  @echo "  just migration-down"
  @echo "  just migration-status"
  @echo "  just sqlc-gen"
  @echo "  just buf-gen"
  @echo "  just run"

migration-create name:
  goose -dir {{MIGRATIONS_DIR}} create {{name}} sql

migration-up:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} {{DB_STRING}} up

migration-down:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} {{DB_STRING}} down

migration-status:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} {{DB_STRING}} status

run:
  go run cmd/server/server.go

sqlc-gen:
  sqlc generate

buf-gen:
  buf generate

mock-gen:
  mockery

test:
  go test -v ./...

gen:
  just buf-gen
  just migration-up
  just sqlc-gen
  just mock-gen

all:
  just buf-gen
  just migration-up
  just sqlc-gen
  just mock-gen
  just run

