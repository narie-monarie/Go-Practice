package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	lisa := user{
		name: "Lisa",
		age:  21,
	}

	var u user
	u.name = "Mask"
	u.age = 21

	n := new(user)
	n.name = "Metropolis"
	n.age = 24

	fmt.Println(lisa.age)
	fmt.Println(u.name, n.name)
}
