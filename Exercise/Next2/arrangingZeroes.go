package main

import "fmt"

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}

	return nums
}

func main() {
	x := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(x))
}
