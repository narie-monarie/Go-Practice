package main

import "fmt"

func sort(arr []int) {
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		minValue := arr[minIndex]
		for j := i; j < len(arr)-1; j++ {
			if minValue > arr[j+1] {
				minValue = arr[j+1]
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

	scores := []int{20, 10, 4, 5, 56, 23, 12}
	sort(scores)
	for i := range scores {

		fmt.Println(scores[i])

	}
}
