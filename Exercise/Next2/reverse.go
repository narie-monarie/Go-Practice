package main

import (
	"fmt"
	"strings"
)

func reverse(w string) string {
	var res strings.Builder
	for i := len(w) - 1; i >= 0; i-- {
		res.WriteByte(w[i])
	}

	return res.String()
}

func main() {
	fmt.Println(reverse("Hello World"))
}
