package main

import "fmt"

func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i, j := range arr {
		if j, ok := m[target-j]; ok {
			return []int{j, i}
		}
		m[j] = i
	}
	return []int{}
}

func main() {
	arr := []int{3, 4, 2, 1, 9, 6}
	target := 7
	fmt.Println(twoSum(arr, target))
}
