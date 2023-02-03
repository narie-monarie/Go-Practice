package main

import "fmt"

func main() {
	var arr []int32
	var n, m, a, b, size, sizeMain int32

	fmt.Scanf("%d", &size)

	sizeMain = size

	for size != 0 {
		fmt.Scan(&n)
		arr = append(arr, n)
		size--
	}

	fmt.Println(sizeMain)
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			a++
		} else if arr[i] < 0 {
			b++
		} else if arr[i] == 0 {
			m++
		}
	}

	fmt.Printf("%f \n", float64(a)/float64(sizeMain))
	fmt.Printf("%f \n", float64(b)/float64(sizeMain))
	fmt.Printf("%f \n", float64(m)/float64(sizeMain))
}
