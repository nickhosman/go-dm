package main

import (
	"database/sql"
	"fmt"
	"math"
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
	Class Class
}

type Message struct {
	Content string `json:"message"`
}

func newMessage(content string) Message {
	return Message{
		Content: content,
	}
}

func newCharacter(id int, name string, lvl int, class Class, player User) Character {
	c := Character{
		Id: id,
		Name: name,
		Lvl: lvl,
		Player: player,
		Class: class,
	}

	return c
}

func GetCharHealth(c Character) int {
	die := c.Class.Hitdie

	hp := die + (int(math.Ceil(float64(die / 2))) * c.Lvl)

	return hp
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

func newClass(id int, name string, hitdie int) Class {
	return Class {
		Id: id,
		Name: name,
		Hitdie: hitdie,
	}
}

func main() {
	const file string = "dm.db"

	db, err := openDB(file)
	if err != nil {
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())

	u := newUser(1, "knocknix")
	cl := newClass(1, "Rogue", 8)
	char := newCharacter(1, "Drig", 1, cl, u)

	health := GetCharHealth(char)
	fmt.Println("HEALTH:", health)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/character", func(c echo.Context) error {
		return c.JSON(http.StatusOK, char)
	})

	e.GET("/features/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		stmt := `SELECT id, name, description, lvl FROM features WHERE id = ` + strconv.Itoa(id)

		f := db.QueryRow(stmt)

		err = f.Err()
		if err != nil {
			return c.JSON(http.StatusNotFound, newMessage("Feature not found."))
		}

		feature := Feature{}
		err = f.Scan(&feature.Id, &feature.Name, &feature.Description, &feature.Lvl)
		if err != nil {
			return c.JSON(http.StatusNotFound, newMessage("Feature not found."))
		}

		return c.JSON(http.StatusOK, feature)
	})

	e.Logger.Fatal(e.Start(":4000"))
}

func openDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// !reorganize features
