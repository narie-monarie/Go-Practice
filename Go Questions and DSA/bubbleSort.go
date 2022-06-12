package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	scores := []int{60, 50, 95, 80, 70}
	sort(scores)

	for i := range scores {
		fmt.Println(scores[i])
	}
}
