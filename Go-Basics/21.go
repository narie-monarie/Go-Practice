package main

import "fmt"

func maximum(arr []int) int {
	for _, i := range arr {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}

	maxValue := arr[len(arr)-1]
	return maxValue
}

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	maxval := maximum(scores)
	fmt.Println(maxval)
}
