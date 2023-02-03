package main

import (
	"fmt"
	"regexp"
)

func main() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?", reStr)
	fmt.Println(match)

}
