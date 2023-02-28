package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y int
	fmt.Scan(&x)
	var a []int

	for x != 0 {
		fmt.Scan(&y)
		a = append(a, y)
		x--
	}

	var n, m int
	fmt.Scan(&n)
	sort.Ints(a)
	for n != 0 {
		fmt.Scan(&m)
	}

}
