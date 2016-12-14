package main

import (
	"fmt"
)

func main() {
	cities := map[string]int{
		"New York":    123123123,
		"Los Angeles": 23565656,
		"Chicago":     1234534567,
	}
	for key, value := range cities {
		fmt.Printf("%s had %d inhabitans\n", key, value)
	}
}
