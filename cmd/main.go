package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id 		 int `json:"id"`
	Username string `json:"username"`
}

func newUser(id int, username string) User {
	return User{
		Id: id,
		Username: username,
	}
}

type Character struct {
	Id 	   int `json:"id"`
	Name   string `json:"name"`
	Player User `json:"player"`
}

func newCharacter(id int, name string, player User) Character {
	return Character{
		Id: id,
		Name: name,
		Player: player,
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	u := newUser(1, "knocknix")
	char := newCharacter(1, "Drig", u)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/character", func(c echo.Context) error {
		return c.JSON(http.StatusOK, char)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
