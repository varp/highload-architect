package main

import (
	"fmt"
	"os"

	openApiMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/ports/http/v1"
	userRepo "go.vardan.dev/highload-architect/social-net-01/internal/domain/repo/user"
	userUsecase "go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases/user"
	"go.vardan.dev/highload-architect/social-net-01/pkg/tools"
)

func main() {
	uu := userUsecase.New(userRepo.New(), tools.NewUUIDGenerator(), tools.NewPasswordEncryptor())
	apiV1 := v1.NewApi(uu)

	swagger, err := v1.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	g := gin.Default()
	g.Use(openApiMiddleware.OapiRequestValidator(swagger))

	v1.RegisterHandlers(g, apiV1)

	g.Run(":8000")
}
