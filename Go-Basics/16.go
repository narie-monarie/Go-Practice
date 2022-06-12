package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60}
	newslice := slice[0:5]
	newslice[2] = 14
	newslice = append(newslice, 41)
	//SLICING THE LENGTH AND THE CAPACITY
	source := newslice[1:4:4]

	fmt.Println(newslice)
	fmt.Println(source)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("the array is %v", append(s1, s2...))
}
