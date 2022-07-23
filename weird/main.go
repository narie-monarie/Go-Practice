package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	var d []int
	x := 0
	fmt.Scanf("%d", &a)
	for i := 0; i < a; i++ {
		fmt.Scanf("%d", b)
		for j := 0; j < b; j++ {
			fmt.Scanf("%d", &c)
			d = append(d, c)
			sort.Ints(d)
			x = d[i] + d[i+1]

			fmt.Println(j)
		}
	}
}
