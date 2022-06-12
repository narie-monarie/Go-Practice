package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for i := 0; i < a+1; i++ {
		x := strings.Repeat("#", i)
		fmt.Println(x)
	}

}
