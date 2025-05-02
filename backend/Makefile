include .env

GOOSE_DBSTRING ?= $(STR_POSTGRESQL)
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= postgres

APP_NAME := server

air:
	air --build.cmd "go build -o ./tmp/server.exe ./cmd/server/main.go" --build.bin "D:\study\DACN\github.com\QuanCters\backend\tmp\server.exe"

dev:
	go run ./cmd/$(APP_NAME)

up_by_one:
	@set GOOSE_DRIVER=$(GOOSE_DRIVER) && set GOOSE_DBSTRING=$(GOOSE_DBSTRING) && goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	goose $(GOOSE_DRIVER) -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DBSTRING) up

downse:
	goose $(GOOSE_DRIVER) -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DBSTRING) down

resetse:
	goose $(GOOSE_DRIVER) -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DBSTRING) reset

docker_up: 
	docker compose up -d 

docker_down: 
	docker compose down

sqlgen:
	sqlc generate

.PHONY: dev up_by_one create_migration upse downse resetse docker_up docker_down sqlgen