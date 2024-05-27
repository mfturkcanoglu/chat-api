package server

import (
	"chat-api/internal/config"
	"github.com/labstack/echo/v4"
)

var (
	e      *echo.Echo     = echo.New()
	logger echo.Logger    = e.Logger
	c      *config.Config = config.NewConfig(logger)
)

func Serve() {
	logger.Fatal(e.Start(c.Port))
}
