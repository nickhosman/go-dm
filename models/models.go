package models

import "math"

type User struct {
	Id 		 int `json:"id"`
	Username string `json:"username"`
}

func NewUser(id int, username string) User {
	return User{
		Id: id,
		Username: username,
	}
}

type Stat struct {
	Id   int
	Name string
	Val  int
}

func (s *Stat) GetMod() int {
	return int(math.Floor(float64((s.Val - 10) / 2)))
}

func NewStat(id int, name string, val int) Stat {
	return Stat{
		Id: id,
		Name: name,
		Val: val,
	}
}

type Stats struct {
	Str Stat
	Dex Stat
	Con Stat
	Int Stat
	Wis Stat
	Cha Stat
}

func NewStats(str, dex, con, intel, wis, cha Stat) Stats {
	return Stats{
		Str: str,
		Dex: dex,
		Con: con,
		Int: intel,
		Wis: wis,
		Cha: cha,
	}
}

type Feature struct {
	Id 			int
	Name 		string
	Description string
	Lvl 		int
}

func NewFeature(id int, name, description string, lvl int) Feature {
	return Feature{
		Id: id,
		Name: name,
		Description: description,
		Lvl: lvl,
	}
}

type Class struct {
	Id 		int
	Name 	string
	Hitdie  int
	Lvl 	int
}

type Character struct {
	Id 		int
	Name 	string
	Lvl 	int
	Stats 	Stats
	Classes []Class
	Player 	User
}
