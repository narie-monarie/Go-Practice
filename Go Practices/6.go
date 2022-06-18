package main

import "fmt"

func main() {
	x := []int{2, 4, 5, 2, 3, 5, 7}
	for i := len(x) - 1; i >= 0; i-- {
		fmt.Printf("%d ", x[i])
	}
}
