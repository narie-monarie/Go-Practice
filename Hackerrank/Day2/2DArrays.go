package main

import "fmt"

func main() {
	var arr [6][6]int

	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	var s, m int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s = (arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2])
			if s > m {
				m = s
			}
		}
	}

	fmt.Println(" ", m)
}
