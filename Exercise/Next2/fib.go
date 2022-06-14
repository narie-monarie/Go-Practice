package main

import "fmt"

//memoization
//reminder hash maps
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	var n2, n1 int = 0, 1

	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}

func main() {
	fmt.Println(fib(50))
}
