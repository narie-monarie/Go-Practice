package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}

func main() {
	scores := []int{3, 5, -4, 8, 11, 1, -1, 6}
	x := twoSum(scores, 11)
	fmt.Println(x)
}
