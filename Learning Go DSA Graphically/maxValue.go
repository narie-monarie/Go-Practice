//getting the maximum value in an array
package main

import "fmt"

func getMax(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	maxValue := arr[len(arr)-1]
	return maxValue
}

func main() {
	scores := []int{60, 40, 20, 10, 11, 65, 13}
	fmt.Println(getMax(scores))
}
