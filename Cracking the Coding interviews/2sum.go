package main

import "fmt"

func Twosum(arr []int, target int) []int {
	m := make(map[int]int)
	for j, i := range arr {
		if i, ok := m[target-i]; ok {
			return []int{i, j}
		}
		m[i] = j
	}
	return []int{}
}
func main() {
	scores := []int{1, 4, 3, 6, 3, 3, 8, 9, 6, 5, 3}
	fmt.Println(Twosum(scores, 5))
}
