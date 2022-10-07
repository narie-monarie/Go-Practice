package main

import (
	"fmt"
	"math"
)

func max(arr []int, n int) int {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	var maxValue = arr[n-1]
	return maxValue
}

func max2(arr []float64) int {
	var temp float64

	for i := 0; i < len(arr)-1; i++ {
		temp = float64(math.Max(arr[i], temp))
	}

	return int(temp)
}
func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var maxValue = max(scores, length)
	fmt.Printf("Max Value = %d\n", maxValue)
}
