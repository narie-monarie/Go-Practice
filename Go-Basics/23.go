//select sorting

package main

import "fmt"

func sort(arr []int) {
	var minIndex int

	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		minvalue := arr[minIndex]
		for j := i; j < len(arr)-1; j++ {
			if minvalue > arr[j+1] {
				minvalue = arr[j+1]
				minIndex = j + 1
			}
		}
		if i != minIndex {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}
}

func main() {
	scores := []int{20, 13, 12, 21, 33, 11, 4, 6, 19}
	sort(scores)

	for i := range scores {
		fmt.Println(scores[i])
	}
}
