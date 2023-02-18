.PHONY: go \
tools-oapi-codegen generate-server-and-types

go:
	@command -v go >/dev/null && go version

tools-oapi-codegen: go
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-server-and-types: tools-oapi-codegen
	oapi-codegen -package v1 -generate server,spec -o internal/domain/ports/http/v1/server.gen.go api/openapi.yaml
	oapi-codegen -package v1 -generate types -o internal/domain/ports/http/v1/types.gen.go api/openapi.yaml

serve: go
	go run ./cmd/server