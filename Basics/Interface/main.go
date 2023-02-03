package main

import "fmt"

type Animal interface {
	angrySound()
	happySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("Cat Attack!")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) angrySound() {
	fmt.Println("Cat Hissss")
}

func (c Cat) happySound() {
	fmt.Println("Cat Purr!!")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.angrySound()
}
