package main

import "fmt"

func permutation(sampleRune []rune, left, right int) {
	if left == right {
		fmt.Println(string(sampleRune))
	} else {
		for i := left; i <= right; i++ {
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			permutation(sampleRune, left+1, right)
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
}

func main() {
	sample := "abc"
	sampleRune := []rune(sample)
	permutation(sampleRune, 0, len(sampleRune)-1)
}
