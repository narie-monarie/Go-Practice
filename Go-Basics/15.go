package main

import "fmt"

func main() {
	// a slice is a data structure that provides a way
	//for one to work with and manage collections of data
	//Slices are tiny objects that abstract and manipulate
	//an underlying array

	//slice := make([]string, 5)
	//has a capacity of 5
	//slice = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//slice2 := make([]int, 3, 5) //has length 3
	//capacity of 5

	//nil slices have nil pointer, nil length and nil capacity

	slice3 := []int{1, 2, 3, 4}
	slice3[1] = 4
	fmt.Println(slice3)

	slice4 := []int{10, 20, 30, 40, 50}
	slice5 := slice4[1:3]
	slice5[1] = 33
	fmt.Println(slice5)
}
