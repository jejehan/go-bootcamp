package main

//Basic method
import (
	"fmt"
)

type User struct {
	Firstname, Lastname string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Hallo %s %s", u.Firstname, u.Lastname)
}

func main() {
	user := User{"Jason", "Bourne"}
	fmt.Println(user.Greeting())
}
