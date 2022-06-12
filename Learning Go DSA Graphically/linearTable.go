//Linear - Sequence of elements in a one-dimensional array
package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}

	for _, i := range scores {
		fmt.Printf("%d,", i)
	}
}
