package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Content string
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return(c.JSON(http.StatusOK, Message{"Hello world."}))
	})

	e.Logger.Fatal(e.Start(":4000"))
}
