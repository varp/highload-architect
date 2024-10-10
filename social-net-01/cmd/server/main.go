package main

import (
	"fmt"
	"os"

	openApiMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"

	"go.vardan.dev/highload-architect/social-net-01/internal/domain/ports/http/v1"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases"
	"go.vardan.dev/highload-architect/social-net-01/pkg/tools"
)

func main() {
	const (
		errLoadingSwagger = iota + 1
	)

	uu := usecases.NewUserUsecase(repos.NewUserRepository(), tools.NewPasswordEncryptor())
	apiV1 := v1.NewApi(uu)

	swagger, err := v1.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(errLoadingSwagger)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := gin.Default()
	r.Use(openApiMiddleware.OapiRequestValidator(swagger))
	v1.RegisterHandlers(r, apiV1)

	r.Run(":8000")
}
