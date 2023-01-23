package main

import "fmt"

func main() {
	var n, i, num, one, two, indOne, indTwo int64
	var arr []int64
	fmt.Scan(&num)
	for num != 0 {
		fmt.Scan(&n)
		arr = append(arr, n)
		num--
	}

	for i = 0; i < int64(len(arr)); i++ {
		if arr[i]%2 == 0 {
			two++
			indTwo = i
		} else {
			one++
			indOne = i
		}
	}

	if one > two {
		fmt.Print(indTwo + 1)
	} else {
		fmt.Print(indOne + 1)
	}
}
