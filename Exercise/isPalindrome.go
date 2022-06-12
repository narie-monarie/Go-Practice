package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	lp, rp := 0, len(s)-1
	for lp <= rp {
		if s[lp] != s[rp] {
			return false
		}
		lp++
		rp--
	}
	return true
}

func main() {
	fmt.Println("Hello World")
}
