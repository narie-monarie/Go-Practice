//linear Tables
package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}

	length := len(scores)
	tempArray := make([]int, length+1) //create a new Array with +1 size the original

	for i := 0; i < length; i++ {
		tempArray[i] = scores[i]
	}

	tempArray[length] = 75

	scores = tempArray

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
