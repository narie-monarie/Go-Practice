//Linear Table is a sequence of elements in a one dimension array
package main

import "fmt"

func main() {
	scores := []int{10, 24, 32, 11, 43, 5}
	length := len(scores)
	for i := 0; i < length; i++ {
		fmt.Println(scores[i])
	}
}
