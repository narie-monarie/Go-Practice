package main

import "fmt"

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	if k <= (n+1)/2 {
		fmt.Println(k*2 - 1)
	} else {
		fmt.Println((k - (n+1)/2) * 2)
	}

}
