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
	Id             int
	Name, Location string
	GameId         int
}

func main() {
	p := Player{}
	p.Id = 12
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 1232
	fmt.Printf("%+v \n", p)
}
