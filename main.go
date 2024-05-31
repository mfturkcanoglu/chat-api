package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type Messagebody struct {
	Message string `json:"message"`
}

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	welcomeMessage := &Messagebody{
		Message: "Hello there!",
	}
	err = ws.WriteJSON(welcomeMessage)
	if err != nil {
		c.Logger().Error(err)
	}

	for {
		var msg Messagebody
		err = ws.ReadJSON(&msg)

		if err != nil {
			c.Logger().Error(err)
		}

		c.Logger().Print(msg)

	}
}

func main() {

	//server.Serve()
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%s: uri: %v, status: %v\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":8080"))
}
