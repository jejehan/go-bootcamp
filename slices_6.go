package main

import (
	"fmt"
)

func main() {
	cities := []string{
		"Santa Monica",
		"Jakarta",
		"Bandung",
	}
	fmt.Println(len(cities))
	countries := make([]string, 42)
	fmt.Printf("%q\n", cities)
	fmt.Printf("%q\n", countries)
	fmt.Println(len(countries))
}
