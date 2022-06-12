package main

import "fmt"

func returner() (int, string) {
	i := "Hello"
	x := 1
	return x, i
}

func main() {
	a, b := returner()
	fmt.Println(a, b)
}
