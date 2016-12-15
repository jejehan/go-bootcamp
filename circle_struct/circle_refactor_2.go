package main

//refactor using method

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) Area() float64 {
	return math.Pi * (c.r * c.r)
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.Area())
}
