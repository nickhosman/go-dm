package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Username string `json:"username"`
}

func newUser(username string) User {
	return User{
		Username: username,
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	u := newUser("knocknix")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
