package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	scores := []int{20, 40, 31, 45, 76, 34, 8, 98, 3}

	bubbleSort(scores)
	for _, i := range scores {
		fmt.Printf("%d,", i)
	}
}
