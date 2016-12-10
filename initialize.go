package main

import (
	"fmt"
)

type Bootcamp struct {
	Lat float32
	Lon float32
}

func main() {
	x := new(Bootcamp)
	x.Lat = 12.3123
	y := Bootcamp{}
	y.Lat = 12.3123
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*x == y)
}
