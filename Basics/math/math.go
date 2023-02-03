package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	seedValue := time.Now().Unix()
	rand.Seed(seedValue)
	pl(seedValue)
	pl("5 + 4", 5+4)
	x := math.Abs(-10)
	pl(x)
}
