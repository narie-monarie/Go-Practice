package main

import (
	"fmt"
	"math"
)

func main() {
	var a string
	var b, m, c, n float64
	fmt.Scan(&a)

	for i := 0; i < len(a); i++ {
		if a[i] == '0' {
			b++
		} else if a[i] == '1' {
			b = 0
		}
		m = math.Max(m, b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			c++
		} else if a[i] != '1' {
			c = 0
		}
		n = math.Max(n, c)
	}

	//fmt.Println(m, n)

	if m >= float64(7) {
		fmt.Println("YES")
	} else if n >= float64(7) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
