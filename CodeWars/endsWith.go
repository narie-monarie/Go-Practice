package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	if strings.Contains(str, ending) {
		return true
	}
	return false
}

//alternative Solutions

func solutions(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}
func main() {
	fmt.Println(solution("abc", "d"))
}
