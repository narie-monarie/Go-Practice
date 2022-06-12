package main

import "fmt"

func main() {
	names := [3]string{}
	names[0] = "monari"
	names[1] = "Job"
	names[2] = "monari"
	fmt.Println(names)

	//multidirectional Arrays

	coordinates := [3][3]int{}
	coordinates[1][2] = 10
	fmt.Println(coordinates)
}
