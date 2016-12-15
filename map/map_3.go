package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float32
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	delete(m, "Google")
	fmt.Println(m)
}
