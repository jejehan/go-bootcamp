package main

import (
	"fmt"
	"time"
)

type Bootcamp struct(
    Lat, Lon float32
    Date time.Time
)

func main() {
	event := Bootcamp{
		Lat: 34.09293993,
		Lon: -123.9292,
	}
	event.Date = time.Now
	fmt.Printf("Event on %s, location (%f, %f)", event.Date, event.Lan, event.Lon)
}
