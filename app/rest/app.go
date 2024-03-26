package main

import (
	"github.com/labstack/echo/v4"
	"sinarmas/app"
	"sinarmas/internal/handler/hello"
	"sinarmas/internal/model/config"
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

	return e, nil
}
