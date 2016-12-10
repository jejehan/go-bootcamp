package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

func main() {
	fmt.Println(Bootcamp{
		Lat:  34.09299828,
		Lon:  -118.9828829,
		Date: time.Now(),
	})
}
