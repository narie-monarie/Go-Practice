package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(x)-1; i++ {
		if i%2 == 0 {
			var temp = x[i]
			x[i] = x[i+1]
			x[i+1] = temp
		}
	}

	for _, i := range x {
		fmt.Printf("%d ", i)
	}
}
