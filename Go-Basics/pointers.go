package main

import "fmt"

//pointers - marked by the [&] sign store only the memory address of a variable
func main() {
	hello := "mine"
	mine := &hello // pointers
	*mine = "maniac"
	fmt.Println(*mine)
}
