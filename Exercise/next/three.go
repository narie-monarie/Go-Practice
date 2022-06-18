package main

import (
	"fmt"
	"math"
)

func three(arr []int, target int) int {
	indexHigh := len(arr) - 1
	indexLow := 0

	for indexLow <= indexHigh {
		indexM := math.Floor(float64(indexHigh+indexLow) / 2)
		indexMid := int(indexM)
		if arr[indexMid] == target {
			return indexMid
		} else if arr[indexMid] < target {
			indexLow = indexMid + 1
		} else {
			indexHigh = indexMid - 1
		}
	}

	return -1
}

func main() {
	vamos := []int{11, 12, 34, 45, 65, 87, 98, 564}
	fmt.Println(three(vamos, 87))
}
