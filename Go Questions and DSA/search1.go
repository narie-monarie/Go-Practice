//linear search unsorted input
package main

import "fmt"

func LinearUnsorted(arr []int, val int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if val == data[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
