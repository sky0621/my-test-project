# ###########################################################
# env
# ###########################################################
OPENAPI_SCHEMA_DIR := ./schema/openapi
OPENAPI_PARTS_DIR := $(OPENAPI_SCHEMA_DIR)/parts
OPENAPI_PARTS_ROOT_YAML := $(OPENAPI_PARTS_DIR)/root.yaml
OPENAPI_SCHEMA_YAML := $(OPENAPI_SCHEMA_DIR)/api.yaml

# ###########################################################
# Open API
# ###########################################################
.PHONY: redoc-lint
redoc-lint:
	npx @redocly/cli lint $(OPENAPI_PARTS_ROOT_YAML)

.PHONY: redoc
redoc: redoc-lint
	npx @redocly/cli bundle $(OPENAPI_PARTS_ROOT_YAML) -o $(OPENAPI_SCHEMA_YAML)

.PHONY: generate-api
generate-api: redoc
	go tool oapi-codegen --config=oapi-codegen-config.yaml -o backend/internal/api/generated.go $(OPENAPI_SCHEMA_YAML)

# ###########################################################
# SQLC
# ###########################################################
.PHONY: generate-sqlc
generate-sqlc:
	go tool sqlc generate

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
.PHONY: build
build:
	mkdir -p bin
	go build -o bin/server ./backend/cmd/server

.PHONY: run
run: build
	./bin/server
