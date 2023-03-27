package main

import "fmt"

func main() {
	x := make(map[int]int)
	tar := []int{1, 2, 3, 4, 5}
	target := 8
	for i, v := range tar {
		if b, ok := x[target-v]; ok {
			fmt.Println([]int{tar[b], tar[i]})
		}
		x[v] = i
	}
}
