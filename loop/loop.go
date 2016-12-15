package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := float64(1)
	z = 1.0

	aproximation := func(z, x float64) float64 {
		return z - ((z*z)-x)/(2*z)
	}
	x = 1.0
	for i := 0; i < 10; i++ {
		z = aproximation(z, x)
	}
	return z
}

func main() {
	for i := 1.0; i < 11.0; i++ {
		fmt.Printf("%d: %f vs %f\n", int(i), Sqrt(i), math.Sqrt(i))
	}
}
