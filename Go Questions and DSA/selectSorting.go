//select sort an array in GO

package main

import "fmt"

func selectSort(arr []int) {
	min_index := 0
	for i := 0; i < len(arr)-1; i++ {
		min_index = i
		min_value := arr[min_index]
		for j := i; j < len(arr)-1; j++ {
			if min_value > arr[j+1] {
				min_value = arr[j+1]
				min_index = j + 1
			}
		}
		if i != min_index {
			temp := arr[i]
			arr[i] = arr[min_index]
			arr[min_index] = temp
		}
	}
}

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	selectSort(scores)

	for i := range scores {
		fmt.Println(scores[i])
	}
}
