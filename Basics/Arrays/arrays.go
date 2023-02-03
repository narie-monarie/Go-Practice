package main

import "fmt"

var pl = fmt.Println

func main() {
	one := "Hello Maniac"
	arrOne := []rune(one)

	for _, i := range arrOne {
		pl(i)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bstr := string(byteArr[:])
	pl(bstr)
}