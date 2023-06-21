.PHONY: go tools-oapi-codegen up

define putenv
	@test -f .env || cp default.env .env
	@sed -e '/$1=/d' .env > .env.bak && mv .env.bak .env
	@echo "$1=$2" >> .env
endef

export DOCKER_BUILDKIT = 1
export PROJECT_NAME = social-net
export APP = $(PROJECT_NAME)-server
export GIT_SHA = $(shell git rev-parse --short HEAD)
export VERSION = $(shell git describe --tags | cut -f2 -d'/')

MIGRATIONS_DIR ?= sql/migrations


# tools
go:
	@command -v go >/dev/null && go version

tools-oapi-codegen: go
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-server-and-types: tools-oapi-codegen
	oapi-codegen -package v1 -generate gin,spec -o internal/domain/ports/http/v1/server.gen.go api/openapi.yaml
	oapi-codegen -package v1 -generate types -o internal/domain/ports/http/v1/types.gen.go api/openapi.yaml

# migrations
sql-create-migration:
	touch "$(MIGRATIONS_DIR)/$(shell date +%Y%m%d%H%M%S).sql"

sql-migrate:
	docker-compose exec db /bin/sh -c "/scripts/migrate.sh"

# deploy
up-infra up-dev: COMPOSE_FILE=deployments/docker-compose.dev.yml
update-env:
	$(call putenv,COMPOSE_FILE,$(COMPOSE_FILE))
	$(call putenv,COMPOSE_PROJECT_NAME,$(PROJECT_NAME))

up-infra: update-env
	docker-compose up -d --remove-orphans db
up-dev: update-env
	docker-compose up -d --remove-orphans app
up: up-infra up-dev

stop:
	docker-compose stop
down:
	docker-compose down

run:
	go run ./cmd/server


# build
@build:
	CGO_ENABLED=0 go build -v -o social-net-server ./cmd/server

@app:
	docker buildx build -f deployments/app/Dockerfile -t $(APP):$(VERSION) -t $(APP):latest .
