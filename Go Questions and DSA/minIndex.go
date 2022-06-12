package main

import "fmt"

func min(arr []int) int {
	min_value := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[min_value] > arr[i] {
			min_value = i
		}
	}

	return arr[min_value]

}

func main() {
	scores := []int{60, 80, 95, 50, 70}
	minValue := min(scores)
	fmt.Printf("Min Value = %d\n", minValue)
}
