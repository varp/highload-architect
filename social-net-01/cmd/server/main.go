package main

import (
	"fmt"
	"os"

	opapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/ports/http/v1"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repo"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases"
)

func main() {
	userUsecase := usecases.NewUserUsecase(repo.NewUserRepository())
	apiV1 := v1.NewApi(userUsecase)

	swagger, err := v1.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(opapiMiddleware.OapiRequestValidator(swagger))

	v1.RegisterHandlers(e, apiV1)

	e.Logger.Fatal(e.Start(":8000"))
}
