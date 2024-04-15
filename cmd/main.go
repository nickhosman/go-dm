package main

import (
	"net/http"
	"strconv"

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

type Message struct {
	Content string `json:"message"`
}

func newMessage(content string) Message {
	return Message{
		Content: content,
	}
}

var Characters = make(map[int]Character)

func newCharacter(id int, name string, player User) Character {
	c := Character{
		Id: id,
		Name: name,
		Player: player,
	}

	Characters[c.Id] = c

	return c
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

	e.GET("/character/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		if Characters[id].Id > 0 {
			return c.JSON(http.StatusOK, Characters[id])
		}

		return c.JSON(http.StatusNotFound, newMessage("Character not found."))
	})

	e.Logger.Fatal(e.Start(":4000"))
}
