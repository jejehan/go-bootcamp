package main

import (
	"fmt"
)

func main() {
	a := [2]string{"hallo", "world"}
	fmt.Println(a)
	fmt.Println("%s\n", a)
	fmt.Println("%q\n", a)
}
