package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Hello There mate"
	v := strings.Split(x, " ")
	for i := 0; i < len(v); i++ {
		for j := len(v) - i; j >= i; j-- {
			fmt.Printf("%s ", v[i])
		}
	}
}
