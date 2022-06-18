package main

import (
	"fmt"
	"regexp"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func main() {
	letters := "Alex||!!"
	fmt.Println(IsLetter(letters))
}
