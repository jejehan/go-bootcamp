package main

import (
	"fmt"
)

func main() {
	cities := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Jakarta", "bandung"}
	cities = append(cities, otherCities...)
	fmt.Printf("%q\n", cities)
}
