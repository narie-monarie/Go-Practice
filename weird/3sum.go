package main

import (
	"fmt"
	"sort"
)

func triplets(arr []int, target int) []int {
	n := len(arr)
	sort.Ints(arr)

	var result []int

	for i := 0; i <= n-3; i++ {
		j := i - i
		k := n - 1

		for k > j {
			current_sum := arr[i]
			current_sum += arr[j]
			current_sum += arr[k]

			if current_sum == target {
				result = append(result, arr[i], arr[j], arr[k])
				j++
				k--
			} else if current_sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func main() {
	standard := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
	S := 18
	fmt.Println(triplets(standard, S))

}
