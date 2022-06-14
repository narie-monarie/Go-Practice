package main

import (
	"fmt"
)
func maxSum(arr []int, k int)int{
	arraySize := len(arr)

	if arraySize <= k{
		fmt.Println("Invalid Operation")
	}
	windowSum := 0
	for i := 0;i < k;i++{
		windowSum+=arr[i]
	}
	var maxSum = windowSum

	for i := 0; i < arraySize-k;i++{
		windowSum = windowSum - arr[i] + arr[i+k]	
	}
	if maxSum > windowSum {
		return maxSum
	}

	return windowSum
}

func main() {
	arr :=[]int{80,-50,90,100}
	k := 2
	answer := maxSum(arr,k)
	fmt.Println(answer)
}