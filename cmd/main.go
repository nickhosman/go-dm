package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
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
	Lvl    int `json:"lvl"`
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

func newCharacter(id int, name string, lvl int, player User) Character {
	c := Character{
		Id: id,
		Name: name,
		Lvl: lvl,
		Player: player,
	}

	Characters[c.Id] = c

	return c
}

type Feature struct {
	Id			int `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
	Lvl			int `json:"lvl"`
}

func newFeature(name, description string, id, lvl int) Feature {
	return Feature {
		Id: id,
		Name: name,
		Description: description,
		Lvl: lvl,
	}
}

type Class struct {
	Id 	     int `json:"id"`
	Name     string `json:"name"`
	Hitdie   int `json:"hitdie"`
	Lvl 	 int `json:"lvl"`
}

func newClass(id int, name string, hitdie int, features []Feature) Class {
	return Class {
		Id: id,
		Name: name,
		Hitdie: hitdie,
	}
}

func main() {
	startDB()

	e := echo.New()

	e.Use(middleware.Logger())

	u := newUser(1, "knocknix")
	char := newCharacter(1, "Drig", 1, u)

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

func startDB() {
	const file string = "dm.db"

	db, _ := sql.Open("sqlite3", file)

	const create string = `
	CREATE TABLE IF NOT EXISTS features (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		lvl INTEGER
	);`

	db.Exec(create)
}
