package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{
		User{
			Id:       123213,
			Name:     "Matt",
			Location: "LA",
		},
		123213,
	}

	fmt.Printf("Id: %d, Name: %s, Location: %s, GameId: %d \n", p.Id, p.Location, p.Location, p.GameId)
}
