package main

import "fmt"

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Location string
	Box      int
	City     string
}

func main() {

	spiderman := Superhero{
		Name: "Toby Maguire",
		Age:  21,
		Address: Address{
			Location: "Ngong",
			Box:      10245,
			City:     "Nairobi",
		},
	}

	fmt.Println(spiderman.Address.City)
}
