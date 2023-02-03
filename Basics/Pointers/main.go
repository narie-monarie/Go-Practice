package main

import "fmt"

func ChangedVAlue(myPtr *int) {
	*myPtr = 12
}
func dblArrayVAls(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	//pointers - point location of a value inside memory

	f4 := 10
	fmt.Println("f4: ", f4)
	fmt.Println("f4 Address: ", &f4)

	var f4Ptr *int = &f4
	fmt.Println("f4Ptr Address: ", f4Ptr)
	fmt.Println("f4Ptr Address Value: ", *f4Ptr)

	*f4Ptr = 4
	fmt.Println("f4Ptr changed Value: ", *f4Ptr)

	fmt.Println("f4 before function :", f4)
	ChangedVAlue(&f4)
	fmt.Println("new f4 ", f4)

	prr := [4]int{1, 2, 3, 4}
	dblArrayVAls(&prr)
	fmt.Println(prr)
}
