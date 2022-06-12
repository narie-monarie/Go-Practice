/*Go Program to Count Digits in a Number

Write a Go Program to Count Digits in a Number using for loop.
The for loop condition (num > 0) make sure the number is greater than zero. \
Within the loop, we are incrementing the count value.
Next, we divide the number by ten (num = num / 10),
which will remove the last digit from a number.*/

package main

import "fmt"

func main() {
	var num int
	count := 0
	fmt.Scanln(&num)

	for num > 0 {
		num = num / 10
		count++
	}

	fmt.Println(count)
}
