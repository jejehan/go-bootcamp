package main

import (
	"fmt"
)

func main() {
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[2] = "Venice"
	cities[3] = "Los Angeles"
	fmt.Printf("%q\n", cities)
}
