//Get the first digit in a number

package main

import "fmt"

func firstNumber(first int) int {
	for first >= 10 {
		first = first / 10
	}
	return first
}

func main() {
	fmt.Println(firstNumber(320))
}
