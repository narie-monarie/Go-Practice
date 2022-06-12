package main

import "fmt"

func wayTooLongWords() {
	var target int
	fmt.Scanln(&target)
	var arr []string
	var word string
	for i := 0; i < target; i++ {
		fmt.Scanln(&word)
		arr = append(arr, word)
	}
	for _, i := range arr {
		l := 0
		r := len(i) - 1
		if len(i) > 10 {
			fmt.Printf("%c%d%c\n", i[l], len(i)-2, i[r])
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	wayTooLongWords()
}
