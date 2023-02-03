package main

import (
	"fmt"
	"strings"
)

// create custome datatypes
type Customer struct {
	name    string
	address string
	balance float64
}

type Contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	Contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s \n", b.name, b.Contact.fName, b.Contact.phone)
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func getCustomerInfo(c Customer) {
	fmt.Printf("%s owes us %.2f$\n", c.name, c.balance)
}

func newCustomerAddress(c *Customer, address string) {
	c.address = address
}

func main() {
	var tS Customer
	tS.name = "Jerry Smith"
	tS.address = "5 main Street"
	tS.balance = 234.56

	getCustomerInfo(tS)
	newCustomerAddress(&tS, "123 SOuth sTREET")
	println("address: ", strings.ToUpper(tS.address))

	sS := Customer{"Sally Smith", "643 Main Area", 0.03}
	print(sS.name)

	contact1 := Contact{
		"James",
		"Eank",
		"555-23-34",
	}

	bus1 := business{
		"Plumbing",
		"254 Nort Street",
		contact1,
	}

	bus1.info()

}
