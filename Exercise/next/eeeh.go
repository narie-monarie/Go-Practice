package main

import "fmt"

func main() {
	var ar []int64
	var a, b, sum int64
	fmt.Scanln(&a)
	for i := 0; int64(i) < a; i++ {
		fmt.Scanf("%d", &b)
		ar = append(ar, b)
		sum += ar[i]

	}

	fmt.Println(sum)

}
