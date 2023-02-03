package main

import "fmt"

var print = fmt.Println

func main() {
	m := make([]int, 8)

	for i := 0; i < len(m); i++ {
		m = append(m, i)
	}

	for _, i := range m {
		print(i)
	}
}
