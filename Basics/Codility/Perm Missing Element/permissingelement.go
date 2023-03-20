package main

import (
	"fmt"
)

func main() {
	x := [5]int{2, 3, 1, 5}
	//Brute Force
	//sort.Ints(x)
	//for i := range x {
	//	c := i + 1
	//if c != x[i] {
	//		fmt.Println(c)
	//	}
	//}

	var sum int
	for _, v := range x {
		sum += v
	}

	fmt.Println(((len(x) * (len(x) + 1)) / 2) - sum)

}
