package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, a int
	var n, v string
	fmt.Scan(&x)
	m := make(map[string]bool, x)

	for x != 0 {
		fmt.Scan(&n)
		_, ok := m[n]
		a++
		if ok {
			v = n + strconv.Itoa(a)
			fmt.Println(v)
		} else {
			fmt.Println("OK")
		}
		m[n] = true
		x--
	}
}
