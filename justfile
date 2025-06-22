# Variablejust
DB_DRIVER := "postgres"
DB_STRING := "host=localhost user=postgres password=postgres dbname=weather-archiver sslmode=disable"
MIGRATIONS_DIR := "sql/migrations"

setup:
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
  go install github.com/pressly/goose/v3/cmd/goose@latest

migration-create name:
  goose -dir {{MIGRATIONS_DIR}} create {{name}} sql

migration-up:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} {{DB_STRING}} up

migration-down:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} {{DB_STRING}} down

migration-status:
  goose -dir {{MIGRATIONS_DIR}} {{DB_DRIVER}} "{{DB_STRING}}" status

db-run:
  docker compose -f docker-compose.yaml up -d

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

