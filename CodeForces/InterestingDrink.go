package main

import (
	"fmt"
	"sort"
)

func binarySearch(a []int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] <= x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	var x, n int
	fmt.Scan(&x)
	a := make([]int, x)

	for i := 0; i < x; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		m := 0
		fmt.Scan(&m)
		j := binarySearch(a, m)
		fmt.Println(j)
	}

}
