package main

import "fmt"

func main() {
	var a, b, count int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Scanln(&b)
		d := make([]int, b)
		for j := 0; j < b; j++ {
			fmt.Scanf("%d", &c)
			for k := 0; k < len(d)-1; k++ {
				count += d[k]

			}
		}
	}

	if count == 0 {
		fmt.Println("YES")
	} else {

		fmt.Println("NO")
	}
}
