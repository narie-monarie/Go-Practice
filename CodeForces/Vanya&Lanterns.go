package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, v int
	fmt.Scan(&n, &m)
	x := []int{}
	for n != 0 {
		fmt.Scan(&v)
		x = append(x, v)
		n--
	}
	sort.Ints(x)

	for _, i := range x {
		fmt.Printf("%.1d ", i)
	}
	//println(x[len(x)-1])
}
