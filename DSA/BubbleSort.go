package main

import "fmt"

func sort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
