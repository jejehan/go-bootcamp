package main

import "fmt"
import "log"

type User struct {
	Username, Password string
}

func setUser(user *User) {
	user.Username = "A"
	user.Password = "B"
}

func main() {
	xPtr := new(User)
	log.Println(xPtr)
	setUser(xPtr)
	log.Println(xPtr)
	fmt.Println(*xPtr)

	xPtr = &User{"C", "D"}
	setUser(xPtr)
	log.Println(xPtr)
	fmt.Println(*xPtr)

	xPtr = &User{}
	xPtr.Username = "E"
	xPtr.Password = "F"
	setUser(xPtr)
	log.Println(xPtr)
	fmt.Println(*xPtr)
}
