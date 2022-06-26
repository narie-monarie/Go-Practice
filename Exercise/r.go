package main

import "fmt"

var reverse int = 0

func revNumber(palNum int) int {
	var remainder int

	if palNum > 0 {
		remainder = palNum % 10
		reverse = reverse*10 + remainder
		revNumber(palNum / 10)
		return reverse

	} else {
		return 0

	}

}
func main() {

	var palNum int

	fmt.Print("Enter the Number to check Palindrome = ")
	fmt.Scanln(&palNum)

	reverse = revNumber(palNum)
	fmt.Println("The Reverse of the Given Number = ", reverse)

	if palNum == reverse {
		fmt.Println(palNum, " is a Palindrome Number")

	} else {
		fmt.Println(palNum, " is Not a Palindrome Number")

	}

}
