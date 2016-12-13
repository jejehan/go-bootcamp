package main

import (
	"fmt"
)

func main() {
	mySlices := []int{2, 4, 6, 8, 10}
	fmt.Println(mySlices)

	fmt.Println(mySlices[1:4])

	fmt.Println(mySlices[:3])

	fmt.Println(mySlices[4:])
}
