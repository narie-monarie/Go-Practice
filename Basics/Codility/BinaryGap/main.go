package main

import (
	"fmt"
	"math"
	"strconv"
)

func solution(n int64) int64 {
	v := strconv.FormatInt(n, 2)
	var b, maxb int64
	fmt.Println(v)
	for i := 0; i < len(v); i++ {
		if v[i] == '0' {
			b++
		} else if v[i] == '1' {
			maxb = int64(math.Max(float64(maxb), float64(b)))
			b = 0
		}
	}
	return maxb
}
func main() {
	fmt.Println(solution(9))

}
