package main

import "fmt"

type number float64

func returner(n number) number {
	return n
}

func main() {
	fmt.Print(returner(2))
}
