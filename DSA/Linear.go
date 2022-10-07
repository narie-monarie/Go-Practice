package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	for i, _ := range scores {
		fmt.Printf("%d ", scores[i])
	}
}
