package main

import "fmt"

func main() {
	var (
		a = "Monday"
		b = "Tuesday"
	)

	var slct string
	fmt.Scanf("%v", slct)

	switch slct {
	case a:
		fmt.Println("Monday")
	case b:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Please select another Number..")
	}
}
