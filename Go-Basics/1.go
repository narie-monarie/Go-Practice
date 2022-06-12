package main

import "fmt"

func main() {
	var scores = []int{90, 10, 20, 30, 50}

	var length = len(scores)

	for i := 0; i < length; i++ {
		fmt.Printf("%d \n", scores[i])
	}
}
