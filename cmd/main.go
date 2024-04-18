package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/types/descriptorpb"
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
	const file string = "dm.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return
	}

	insert := `INSERT INTO features (name, description, lvl)
	VALUES(?, ?, ?)`
	name := "Ability Score Improvement"
	description := "When you reach 4th level, and again at 8th, 12th, 16th, and 19th level, you can increase one ability score of your choice by 2, or you can increase two ability scores of your choice by 1. As normal, you can't increase an ability score above 20 using this feature."
	lvl := 4

	_, err = db.Exec(insert, name, description, lvl)
	if err != nil {
		return
	}
	
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
