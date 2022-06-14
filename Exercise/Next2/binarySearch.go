package main

import (
	"fmt"
	"math"
)

func BinarySearch(arr []int, target int)int{
	left := 0
	right := len(arr)-1	
	for left <= right{
		var mid1 float64 = float64((left+right)/2)
		mid := int(math.Floor(mid1))
		if arr[mid] == target{
			return mid
		}else if arr[mid] < target{
			left = mid+1
		}else{
			right = mid-1
		}
	}
	
	return -1
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 67
	
	fmt.Println(BinarySearch(arr,target));
}
