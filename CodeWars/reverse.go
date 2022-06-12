package main

import (
	"fmt"
	"strings"
)

func reverse_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	word := strings.Join(words, " ")

	return word
}
func ReverseWord(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	value := string(rns)
	return reverse_words(value)
}

func main() {
	fmt.Println(ReverseWord("This  is  an  Example!"))
}
