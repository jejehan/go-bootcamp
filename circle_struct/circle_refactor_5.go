package main

//refactor using method

import (
	"fmt"
	"math"
)

//interface sama seperti struct bedanya
//isi dari interface adalah method method
//sedangkan struct adalah fields
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

func (c *Circle) perimeter() float64 {
	return 888
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return 666
}

func totalArea(shape ...Shape) float64 {
	var area float64
	for _, s := range shape {
		area += s.area()
	}
	return area
}

//MultiShape memberikan kesempatan untuk mengcompose
//perhitungan yang sama
type MultiShape struct {
	shape []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shape {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shape {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{0, 0, 5}
	//fmt.Println(c.area())
	r := Rectangle{0, 0, 2, 2}
	//fmt.Println(r.area())
	//fmt.Println(r.perimeter())
	//fmt.Println(totalArea(&c, &r))

	m := MultiShape{[]Shape{&c, &r}}
	//fmt.Println(m.area())
	fmt.Println(m.perimeter())
}
