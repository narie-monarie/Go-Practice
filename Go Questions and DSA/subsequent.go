package main

import "fmt"

func subsequence(arr []int, array []int) bool {
	left := 0
	left2 := 0
	for left < len(arr) && left2 < len(array) {
		if arr[left] == array[left2] {
			left2++
		}
		left++
	}
	if left2 == len(array) {
		return true
	}
	return false

}

func main() {
	scores := []int{5, 1, 22, 25, 6, -1, 8, 10}
	scores2 := []int{1, 6, -1, 10}
	fmt.Println(subsequence(scores, scores2))
}
