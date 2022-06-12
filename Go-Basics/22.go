//Bubble Sort
package main

import "fmt"

func sorting(arrays []int) {
	for i := 0; i < len(arrays)-1; i++ {
		for j := 0; j < len(arrays)-1; j++ {
			if arrays[j] > arrays[j+1] {
				temp := arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = temp
			}
		}
	}
}

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	sorting(scores)

	for i := range scores {
		fmt.Println(scores[i])
	}
}
