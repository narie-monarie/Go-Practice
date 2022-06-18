package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	var newNums []int
	left := 0
	for _, i := range nums {
		v := i + nums[left]
		left++
		newNums = append(newNums, v)
	}
	fmt.Println(newNums)
}
