package main

import "fmt"

func arrlist(x []int, target int) []int {
	m := len(x)
	v := make([]int, m)
	for i := 0; i < m; i++ {
		v[(i+target)%m] = x[i]
	}
	return v
}
func main() {
	x := []int{3, 8, 9, 7, 6}
	fmt.Println(arrlist(x, 2))
}
