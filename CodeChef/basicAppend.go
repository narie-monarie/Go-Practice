package main

import "fmt"

func main() {
	var c []int
	a := 4
	var v int
	for a != 0 {
		fmt.Scanf("%d", &v)
		c = append(c, v)
		a--
	}
	c = c[:0]
	for _, i := range c {
		fmt.Println(i)
	}
}
