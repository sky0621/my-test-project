# ###########################################################
# env
# ###########################################################
OPENAPI_SCHEMA_DIR := ./schema/openapi
# for manager
OPENAPI_MANAGER_DIR := $(OPENAPI_SCHEMA_DIR)/manager
OPENAPI_MANAGER_ROOT_YAML := $(OPENAPI_MANAGER_DIR)/root.yaml
OPENAPI_MANAGER_SCHEMA_YAML := $(OPENAPI_SCHEMA_DIR)/manager-api.yaml
# for player
OPENAPI_PLAYER_DIR := $(OPENAPI_SCHEMA_DIR)/player
OPENAPI_PLAYER_ROOT_YAML := $(OPENAPI_PLAYER_DIR)/root.yaml
OPENAPI_PLAYER_SCHEMA_YAML := $(OPENAPI_SCHEMA_DIR)/player-api.yaml

# ###########################################################
# Open API
# ###########################################################
# for manager
.PHONY: redoc-manager-lint
redoc-manager-lint:
	npx @redocly/cli lint $(OPENAPI_MANAGER_ROOT_YAML)

.PHONY: redoc-manager
redoc-manager: redoc-manager-lint
	npx @redocly/cli bundle $(OPENAPI_MANAGER_ROOT_YAML) -o $(OPENAPI_MANAGER_SCHEMA_YAML)

.PHONY: generate-manager-api
generate-manager-api: redoc-manager
	go tool oapi-codegen --config=oapi-codegen-config.yaml -o backend/manager/internal/api/generated.go $(OPENAPI_MANAGER_SCHEMA_YAML)

# for player
.PHONY: redoc-player-lint
redoc-player-lint:
	npx @redocly/cli lint $(OPENAPI_PLAYER_ROOT_YAML)

.PHONY: redoc-player
redoc-player: redoc-player-lint
	npx @redocly/cli bundle $(OPENAPI_PLAYER_ROOT_YAML) -o $(OPENAPI_PLAYER_SCHEMA_YAML)

.PHONY: generate-player-api
generate-player-api: redoc-player
	go tool oapi-codegen --config=oapi-codegen-config.yaml -o backend/player/internal/api/generated.go $(OPENAPI_PLAYER_SCHEMA_YAML)

# ###########################################################
# SQLC
# ###########################################################
.PHONY: generate-sqlc
generate-sqlc: generate-manager-sqlc generate-player-sqlc

.PHONY: generate-manager-sqlc
generate-manager-sqlc:
	go tool sqlc generate -f backend/manager/sqlc.yaml

.PHONY: generate-player-sqlc
generate-player-sqlc:
	go tool sqlc generate -f backend/player/sqlc.yaml

# ###########################################################
# Local DB
# ###########################################################
.PHONY: run-local-mysql
run-local-mysql:
	docker compose up mysql -d --wait

.PHONY: stop-local-mysql
stop-local-mysql:
	docker compose down mysql

.PHONY: run-local-mysql-test
run-local-mysql-test:
	docker compose up mysql_test -d --wait

.PHONY: stop-local-mysql-test
stop-local-mysql-test:
	docker compose down mysql_test

# ###########################################################
# DB Migration
# ###########################################################
.PHONY: migrate-up
migrate-up:
	go run backend/cmd/migration/main.go --up

.PHONY: migrate-down
migrate-down:
	go run backend/cmd/migration/main.go --down

.PHONY: test-migrate-up
test-migrate-up:
	go run backend/cmd/migration/main.go --up --test

.PHONY: test-migrate-down
test-migrate-down:
	go run backend/cmd/migration/main.go --down --test

# ###########################################################
# Test
# ###########################################################
.PHONY: test
test:
	go test ./backend/internal/...

.PHONY: test-short
test-short:
	go test ./backend/internal/... -short

# ###########################################################
# Build & Run Server
# ###########################################################
# for manager
.PHONY: build-manager
build-manager:
	mkdir -p bin
	go build -o bin/manager-server ./backend/cmd/manager/server

.PHONY: run-manager
run-manager: build-manager
	./bin/manager-server

# for player
.PHONY: build-player
build-player:
	mkdir -p bin
	go build -o bin/player-server ./backend/cmd/player/server

.PHONY: run-player
run-player: build-player
	./bin/player-server
