package main

import "fmt"

func Permutation(str []rune, a int) {
	if a == len(str) {
		fmt.Println(string(str))
	} else {
		for i := a; i < len(str); i++ {
			str[a], str[i] = str[i], str[a]
			Permutation(str, a+1)
			str[a], str[i] = str[i], str[a]
		}
	}
}

func main() {
	Permutation([]rune("abc"), 0)
}
