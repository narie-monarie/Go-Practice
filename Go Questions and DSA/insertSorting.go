//Insert sort in Go

package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		insert_element := arr[i]
		insert_position := i
		for j := insert_position - 1; j >= 0; j-- {
			if insert_element < arr[j] {
				arr[j+1] = arr[j]
				insert_position--
			}
		}
		arr[insert_position] = insert_element
	}
}

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	sort(scores)
	for i := range scores {
		fmt.Println(scores[i])
	}
}
