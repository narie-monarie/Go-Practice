package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "lifeJacket", "Paddle"}
	moreNames := [3]string{"Coat", "LifeJacket", "paddle"}
	same := names == moreNames
	fmt.Println(same)
}
