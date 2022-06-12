package main

import "fmt"

func main() {
	var (
		a = 2
		b = 4
	)

	if a > b || b == a {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
