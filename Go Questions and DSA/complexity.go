package main

import "fmt"

//Time Complexity O(n)
func fun1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += 1
	}
	return m
}

//Time complexity 0(n*n)
func main() {
	for i := 0; i < 5; i++ {
		for j := 0; i < 5; j++ {
			fmt.Println(j)
		}
	}
}

//0(log(n))
func fun4(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}

//0(log(n)) - Each time thse space is devided into half eg array/2
func fun5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m += 1
		i = i / 2
	}

	return m
}
