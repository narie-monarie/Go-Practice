package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 6, 8}
	slice = append(slice, 8)
	fmt.Println(cap(slice))

	var twoDArray [8][8]int
	twoDArray[3][6] = 18
	twoDArray[7][4] = 3
	fmt.Println("\n", twoDArray, "\n")
}
