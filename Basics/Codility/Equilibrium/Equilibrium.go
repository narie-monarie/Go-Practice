package main

import "math"

func main() {
	x := []int64{3, 1, 2, 4, 3}
	a := []int64{}
	var sum, n int64
	for i := 0; i < len(x); i++ {
		min := x[i]
		for j := i + 1; j < len(x); j++ {
			sum += x[j]
		}
		n = int64(math.Abs(float64(sum - min)))
		a = append(a, n)
		sum = 0
	}

	for _, l := range a {
		println(l)
	}

}
