.PHONY: go tools-oapi-codegen generate-server-and-types sql-create-migration sql-migrate up-dev up

define putenv
	@touch .env
	@sed -e '/$1=/d' .env > .env.bak && mv .env.bak .env
	@echo "$1=$2" >> .env
endef

export DOCKER_BUILDKIT = 1

PROJECT_DIR="$(shell pwd)"
MIGRATIONS_DIR ?= sql/migrations


# tools
go:
	@command -v go >/dev/null && go version

tools-oapi-codegen: go
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-server-and-types: tools-oapi-codegen
	oapi-codegen -package v1 -generate server,spec -o internal/domain/ports/http/v1/server.gen.go api/openapi.yaml
	oapi-codegen -package v1 -generate types -o internal/domain/ports/http/v1/types.gen.go api/openapi.yaml

# migrations
sql-create-migration:
	touch "$(MIGRATIONS_DIR)/$(shell date +%Y%m%d%H%M%S).sql"

sql-migrate:
	docker-compose exec db /bin/sh -c "/scripts/migrate.sh"

# deploy
up-dev: COMPOSE_FILE=deployments/docker-compose.dev.yml

up-dev:
	$(call putenv,COMPOSE_FILE,$(COMPOSE_FILE))
	docker-compose up -d --remove-orphans

up: up-dev

stop:
	docker-compose stop

serve: go up
	go run ./cmd/server