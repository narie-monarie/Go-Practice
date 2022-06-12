package main

import (
	"fmt"
	"math/rand"
)

//Arrays
func main() {
	var numbers [5]int
	var cities [5]string
	var matrix [3][3]int //2D array
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}
}
