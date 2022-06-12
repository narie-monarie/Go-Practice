package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The area of a Circle")
	var radius float64
	fmt.Scanf("%f ", &radius)

	area := math.Pi * math.Pow(radius, 2)

	fmt.Println("The radius is %.2f and the area  of the Circle is %.2f", radius, area)
}
