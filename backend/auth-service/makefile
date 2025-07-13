include .env

GOOSE_DBSTRING ?= $(STR_POSTGRESQL)
GOOSE_MIGRATION_DIR ?= internal/auth/infrastructure/persistence/schema 
GOOSE_DRIVER ?= postgres

APP_NAME := server

dev: 
	@echo "Starting development server..."
	go run ./cmd/$(APP_NAME)

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	goose $(GOOSE_DRIVER) -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DBSTRING) up

downse:
	goose $(GOOSE_DRIVER) -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DBSTRING) reset

sqlgen:
	sqlc generate

swag:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/server/main.go -o cmd/swag/docs

.PHONY: dev create_migration upse downse sqlgen swag