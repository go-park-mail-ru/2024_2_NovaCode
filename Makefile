ENV_FILE = ./docker/.env
include $(ENV_FILE)

POSTGRES_CONNECTION = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

ENTRY_PATH = ./cmd/main.go
MIGRATIONS_PATH = ./internal/db/$*/migrations
BINARY_PATH = ./bin/main
DOCKER_COMPOSE_PATH = ./docker/docker-compose.yaml
GOLANGCI_LINT_PATH = ./.golangci.yaml

# use `gawk` on mac os
AWK := awk
ifeq ($(shell uname -s), Darwin)
	AWK = gawk
    ifeq (, $(shell which gawk 2> /dev/null))
        $(error "gawk not found")
    endif
endif

################################################################################
# Miscellaneous
################################################################################

.PHONY: help
## (default) Show help page.
help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)";echo;sed -ne"/^## /{h;s/.*//;:d" -e"H;n;s/^## //;td" -e"s/:.*//;G;s/\\n## /---/;s/\\n/ /g;p;}" ${MAKEFILE_LIST}|awk -F --- -v n=$$(tput cols) -v i=29 -v a="$$(tput setaf 6)" -v z="$$(tput sgr0)" '{printf"%s%*s%s ",a,-i,$$1,z;m=split($$2,w," ");l=n-i;for(j=1;j<=m;j++){l-=length(w[j])+1;if(l<= 0){l=n-i-length(w[j])-1;printf"\n%*s ",-i," ";}printf"%s ",w[j];}printf"\n\n";}'

.PHONY: api
## Generate api docs in swagger format.
api:
	@swag init -g $(ENTRY_PATH)

################################################################################
# Local Development
################################################################################

.PHONY: tidy
## Remove unused and add missing dependencies.
tidy:
	@go mod tidy

.PHONY: vendor
## Create `vendor` directory that contains copies of all dependencies
## listed in `go.mod` file.
vendor:
	@go mod vendor

.PHONY: download
## Download modules specified in `go.mod` file.
download:
	@go mod download

.PHONY: build
##@ Build microservice locally.
build: tidy vendor
	@go build -mod vendor -o $(BINARY_PATH) $(ENTRY_PATH)

.PHONY: run
## Run microservice locally.
run:
	@mkdir -p bin
	@go run $(ENTRY_PATH)

.PHONY: exec
## Build microservice locally and run binary.
exec: build
	@$(BINARY_PATH)

################################################################################
# Testing
################################################################################

.PHONY: test
## Run unit tests and generate code coverage report (`./coverage.out`).
test:
	@go clean -testcache
	@go test ./... -coverprofile=coverage.out

.PHONY: --check-coverage
--check-coverage:
	@if [ ! -f coverage.out ]; then \
		echo "coverage does not exist, running tests..."; \
		$(MAKE) test; \
	fi

.PHONY: view-coverage
## View code coverage report if it exists otherwise generate.
view-coverage: --check-coverage
	@go tool cover -func=coverage.out

.PHONY: view-coverage-html
## View code coverage report in browser if it exists otherwise generate.
view-coverage-html: --check-coverage
	@go tool cover -html=coverage.out -o coverage.html
	@open coverage.html

################################################################################
# Formatting & Linting
################################################################################

.PHONY: format
## Format source code.
format:
	@gofmt -s -w .

.PHONY: lint
## Check source code with `golangci-lint` linter.
lint:
	@golangci-lint run --config $(GOLANGCI_LINT_PATH) ./...

################################################################################
# Containers
################################################################################

.PHONY: docker-build
## Build docker container with microservice binary.
docker-build:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_FILE) build

.PHONY: docker-migrate
## Start docker compose service of database and apply migrations.
## Name of docker compose service must match in available database.
## Available databases: postgres, mysql, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `docker-<database>-migrate`.
## Example: `docker-postgres-migrate`.
docker-%-migrate:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_FILE) up -d migrations

.PHONY: docker-start
## Start docker compose containers (all by default).
## Format: `docker-start [compose=<docker-compose-service>]`.
## Example: `docker-start`, `docker-stop compose=postgres`.
docker-start:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_FILE) up -d $(compose)

.PHONY: docker-stop
## Stop docker compose containers (all by default).
## Format: `docker-stop [compose=<docker-compose-service>]`.
## Example: `docker-stop`, `docker-stop compose=postgres`.
docker-stop:
	@docker compose -f $(DOCKER_COMPOSE_PATH) stop $(compose)

.PHONY: docker-ash
## Run `ash` in docker container of microservice.
docker-ash:
	@docker exec -it $(SERVICE_NAME) /bin/ash

.PHONY: docker-psql
## Run `psql` in docker container of postgres.
docker-psql:
	@docker exec -it $(SERVICE_NAME)-postgres psql $(POSTGRES_CONNECTION)

.PHONY: docker-clean
## Remove containers, networks, volumes, and images created by `make docker-start`.
docker-clean:
	@docker compose -f $(DOCKER_COMPOSE_PATH) down

version ?= latest

.PHONY: build-image
## Build docker image of microservice with name.
build-image:
	@docker build -f docker/Dockerfile.$(ENV) --platform linux/amd64 -t daronenko/$(SERVICE_NAME):$(version) .

.PHONY: push-image
## Push docker image of microservice to the docker hub.
push-image:
	@docker push daronenko/$(SERVICE_NAME):$(version)

################################################################################
# Cleaning
################################################################################

.PHONY: clean
## Remove object, cached, and binary files of microservice.
clean:
	@go clean
	@rm -rf $(BINARY_PATH)

.PHONY: clean-test
## Remove test cache and code coverage report.
clean-test:
	@go clean -testcache
	@rm -f coverage.*

.PHONY: clean-all
## Run `make clean` and `make clean-test` commands.
clean-all: clean clean-test

################################################################################
# Database Migrations
################################################################################

.PHONY: %-state
## Show the list of applied migrations to the database.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-state`.
## Example: `postgres-state`.
%-state:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) status

.PHONY: %-migration
## Create migrations with specifed name and type (`sql` - default, `go`).
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-migration name=<name> [type=<sql|go>]`.
## Example: `postgres-migration name=add_some_column type=sql`,
## `postgres-migration name=create_table type=go`.
%-migration:
	$(eval type := $(or $(type), sql))
	@goose -dir $(MIGRATIONS_PATH) $* -s create $(name) $(type)

.PHONY: %-migrate
## Apply all available migrations to the database.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-migrate`.
## Example: `postgres-migrate`.
%-migrate:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) up

.PHONY: %-migrate-to
## Migrate up to a specific version.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-migrate-to version=<version>`.
## Example: `postgres-migrate-to version=20170506082420`.
%-migrate-to:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) up-to $(version)

.PHONY: %-rollback
## Roll back a single migration from the current version.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-rollback`.
## Example: `postgres-rollback`.
%-rollback:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) down	

.PHONY: %-rollback-to
## Roll back migrations to a specific version.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-rollback-to version=20170506082527`.
## Example: `postgres-rollback-to version=20170506082527`.
%-rollback-to:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) down-to $(version)

.PHONY: %-reset
## Roll back all migrations.
## Available databases: postgres, mysql, sqlite3, mssql, redshift, tidb,
## clickhouse, vertica, ydb.
## Format: `<database>-reset`.
## Example: `postgres-reset`.
%-reset:
	$(eval DB := $(shell echo $* | tr '[:lower:]' '[:upper:]'))
	@goose -dir $(MIGRATIONS_PATH) $* $($(DB)_CONNECTION) reset
