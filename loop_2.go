package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {

	aproximation := func(z, x float64) float64 {
		return z - ((z*z)-x)/(2*z)
	}

	i := 0
	z := aproximation(1.0, x)

	for math.Abs(aproximation(z, x)-z) > 0.000001 {
		z = aproximation(z, x)
		i++
	}
	return z, i
}

func main() {
	for i := 1.0; i < 11.0; i++ {
		ours, iterations := Sqrt(i)
		fmt.Printf("%d: in %d iterations %f vs %f\n", int(i), iterations, ours, math.Sqrt(i))
	}
}
