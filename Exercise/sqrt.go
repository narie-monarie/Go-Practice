package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 8
	x := math.Sqrt(a)
	v := math.Floor(x)
	fmt.Println(int(v))
}
