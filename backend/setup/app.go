package setup

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/sky0621/my-test-project/backend/internal/api"
	"github.com/sky0621/my-test-project/backend/internal/shared/config"
	"log"
)

func NewApp() App {
	ctx := context.Background()
	db, err := newDB(ctx, config.NewConfig())
	if err != nil {
		return nil
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	oapiRequestValidator := oapimiddleware.OapiRequestValidator(swagger)

	router := e.Group("/api/v1", oapiRequestValidator)

	api.RegisterHandlers(router, createHandlers(db))

	return &app{echo: e}
}

type App interface {
	Run() error
}

type app struct {
	echo *echo.Echo
}

func (a app) Run() error {
	if err := a.echo.Start(":8080"); err != nil {
		return err
	}
	return nil
}
