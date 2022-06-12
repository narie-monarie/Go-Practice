package main

import "fmt"

func permutation(str []rune, times int) {
	if times == len(str) {
		fmt.Println(string(str))
	} else {
		for i := times; i < len(str); i++ {
			str[times], str[i] = str[i], str[times]
			permutation(str, times+1)
			str[times], str[i] = str[i], str[times]
		}
	}
}

func main() {
	permutation([]rune("Tact Coa"), 0)
}
