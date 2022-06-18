package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	p := r
	nns := make([]int, len(nums))

	for l <= r {
		lv := nums[l] * nums[l]
		rv := nums[r] * nums[r]

		if lv > rv {
			nns[p] = lv
			p -= 1
			l += 1
		} else {
			nns[p] = rv
			p -= 1
			r -= 1
		}
	}

	return nns
}

func main() {
	cores := []int{-7, -5, -4, 3, 6, 8, 9}
	fmt.Println(sortedSquares(cores))
}
