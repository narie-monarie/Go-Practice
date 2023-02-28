package main

import (
	"fmt"
	"sort"
)

func main() {
	var r, x, y, n, m, count int
	fmt.Scan(&x)
	var a []int

	for x != 0 {
		fmt.Scan(&y)
		a = append(a, y)
		x--
	}

	fmt.Scan(&n)
	sort.Ints(a)

	for n != 0 {
		fmt.Scan(&m)

		for i := 0; i < len(a); i++ {
			if m >= a[i] {
				count++
			}
		}
		r = count
		fmt.Println(r)
		count = 0
		n--
	}

}
