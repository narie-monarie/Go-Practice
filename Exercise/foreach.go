package main

import "fmt"

func main() {
	t := "Hello Monari"

	for _, i := range t {
		fmt.Printf("%c ", i)
	}
}
