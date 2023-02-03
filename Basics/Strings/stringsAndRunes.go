package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	sVal := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sVal)
	pl(sV2)
	pl(len(sV2))
	pl(strings.Contains(sV2, "Another"))
	pl(strings.Index(sV2, "o"))
	pl(strings.Replace(sV2, "o", "0", -1))
	pl(strings.Split("12:00AM", "AM"))

}
