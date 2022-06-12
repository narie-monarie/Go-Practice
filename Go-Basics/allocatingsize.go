package main

import "fmt"

func main() {
	names := make([]string, 3, 6)
	//length -> 3
	//capacity -> 6

	names[0] = "Coat"
	names[1] = "LifeJacket"
	names[2] = "Paddle"

	fmt.Println("len: ", len(names))
	fmt.Println("len: ", cap(names))

	moreNames := []string{"Melo", "Menlo"}
	moreNames2 := []string{"Melo", "Menlo"}
	appendNames := append(names, moreNames...)

	//same := moreNames2 == moreNames

	fmt.Println(appendNames)
	//	fmt.Println(same)
}
