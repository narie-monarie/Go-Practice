package main

import "fmt"

func main() {
	//iterating thru arrays

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, value := range slice {
		fmt.Println(value)
	}

	for index := 2; index < len(slice); index++ {
		fmt.Println(slice[index])
	}
}
