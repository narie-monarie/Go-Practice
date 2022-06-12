package main

import "fmt"

func binarySearch(arr []int, val int) bool {
	size := len(arr)
	low := 0
	high := -1
	mid := 0

	//same as while
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] == value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
