package main

import (
	"fmt"
	"math"
	"strconv"
)

func countZeros(N int64) float64 {

	x := strconv.FormatInt(N, 2)

	var a, b float64

	for i := 0; i < len(x); i++ {
		if x[i] == '0' {
			a++
		} else if x[i] == '1' {
			b = math.Max(a, b)
			a = 0
		}
	}
	return b
}

func main() {
	fmt.Println(countZeros(1041))
}
