package main

import "fmt"

func main() {
	arr := []int{4, 3, 5, 1, 2, 6, 9, 2, 10, 11, 12}

	max := arr[0]
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if max > arr[i] {
			max = arr[i]
		} else if min < arr[i] {
			min = arr[i]
		}
	}

	fmt.Printf("%d %d", max, min)
}
