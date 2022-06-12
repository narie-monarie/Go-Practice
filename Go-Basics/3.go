package main

import "fmt"

func main() {
	var str string
	var n, m int
	var nm float32

	str = "Hello World"
	n = 2
	m = 3
	nm = 2.3

	fmt.Printf("value of str = %v \n", str)
	fmt.Printf("2. value = %d \n", n)
	fmt.Printf("3. value = %d \n", m)
	fmt.Printf("3. value = %v \n", nm)

	//Declaring multiple variables
	var (
		name  string
		email string
		age   int
	)

	name = "John"
	email = "John@gmail.com"
	age = 25

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
