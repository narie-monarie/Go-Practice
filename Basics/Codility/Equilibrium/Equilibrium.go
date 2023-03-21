package main

import (
	"fmt"
	"math"
)

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	x := []int{3, 1, 2, 4, 3}

	sum_left := x[0]
	sum_right := Sum(x) - x[0]

	diff := int(math.Abs(float64(sum_left) - float64(sum_right)))

	for i := 1; i < len(x)-1; i++ {

		sum_left += x[i]
		sum_right -= x[i]

		current_diff := int(math.Abs(float64(sum_left) - float64(sum_right)))

		if diff > current_diff {
			diff = current_diff
		}

	}

	fmt.Println(diff)

}
