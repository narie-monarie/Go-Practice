package main

import "fmt"

func rev(arr []int) {
	middle := len(arr) / 2
	for i := 0; i < middle; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}
}

func main() {
	scores := []int{50, 60, 70, 80, 90}
	rev(scores)
	for i := range scores {
		fmt.Println(scores[i])
	}
}
