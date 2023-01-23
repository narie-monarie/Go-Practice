package main

import (
	"fmt"
	"strings"
)

func CapsLock(word string) string {
	if len(word) < 1 {
		return strings.ToLower(word)
	}
	if strings.ToUpper(word) == word {

		return strings.ToLower(word)

	} else if strings.ToLower(word[:1])+strings.ToUpper(word[1:]) == word {

		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])

	}
	return word

}

func main() {
	var word string
	fmt.Scan(&word)
	fmt.Print(CapsLock(word))

}
