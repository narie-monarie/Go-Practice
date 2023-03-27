package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 4, 5, 1, 2, 5, 6, 7}
	mid1 := len(a) / 2
	mid2 := len(a)/2 - 1
	var result, max int
	if ((len(a)/2)-1)%2 != 0 {
		for i := 0; i < len(a)/2; i++ {
			result = a[mid1] - a[mid2]
			mid1++
			mid2--
			max = int(math.Min(float64(result), float64(max)))
		}
	} else {
		for i := 0; i < len(a)/2; i++ {
			mid := len(a) / 2
			result = a[mid1] - a[mid2]
			mid1++
			mid--
			max = int(math.Min(float64(result), float64(max)))
		}

	}

	fmt.Println(max)
}
