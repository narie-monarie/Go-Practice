package main

import "fmt"

func main() {
	var a, b, i float64

	fmt.Scanf("%d %d", &a, &b)
	var initial float64 = 1
	var time float64 = 0

	for i = 0; i < a; i++ {
		var present float64

		fmt.Scan(&present)

		if present >= initial {
			time += present - initial
		} else {
			time += a - (initial - present)
		}
		initial = present
	}

	fmt.Println(time)
}
