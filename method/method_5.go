package main

//Refactor from Basic method at file method.go
import (
	"fmt"
	"log"
)

type User struct {
	Firstname, Lastname string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Hallo %s", n.Name())
}

func main() {
	user := &User{"Jason", "Bourne"}
	user2 := &User{"Matt", "Damon"}
	log.Println(user)
	log.Println(user2)
	fmt.Println(Greet(user))
}
