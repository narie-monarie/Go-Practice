package main

import "fmt"

func main() {
	var n, x int
	var count [5]int
	fmt.Scan(&n)
	for n != 0 {
		fmt.Scan(&x)
		count[x] += 1
		n--
	}

	total := count[4] + count[3] + count[2]/2
	count[1] -= count[3]

	if count[2]%2 == 1 {
		total += 1
		count[1] -= 2
	}
	if count[1] > 0 {
		total += (count[1] + 3) / 4
	}
	fmt.Println(total)
}
