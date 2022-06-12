package main

import "fmt"

func linearSearchSorted(arr []int, value int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
