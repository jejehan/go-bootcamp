package main

import (
	"fmt"
)

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
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
	fmt.Println(p.Greetings())
}
