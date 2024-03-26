package main

import (
	"github.com/labstack/echo/v4"
	"go-be-template/app"
	"go-be-template/internal/handler/hello"
	"go-be-template/internal/middleware/authentication"
	"go-be-template/internal/model/config"
)

func applicationDelegate(cfg *config.Config) (*echo.Echo, error) {
	rsc, err := app.NewCommonResource(cfg)
	if err != nil {
		return nil, err
	}
	repos, err := app.NewCommonRepository(cfg, rsc)
	if err != nil {
		return nil, err
	}
	usecases, err := app.NewCommonUsecase(cfg, repos)
	if err != nil {
		return nil, err
	}
	helloHandler := hello.New(cfg, usecases.HelloUsecase)
	e := echo.New()
	// hello example
	helloGroup := e.Group("/hello")
	helloGroup.GET("", helloHandler.GetHelloMessageHandler)
	helloGroup.GET(
		"/protected",
		helloHandler.GetHelloMessageHandler,
		authentication.AuthorizationMiddleware(cfg.Jwt.AccessSecret))
	return e, nil
}
