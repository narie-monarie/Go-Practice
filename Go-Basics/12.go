// Write a Go Program to Count Digits in a Number using for loop. The for loop condition (num > 0) make sure the number is greater than zero. Within the loop, we are incrementing the count value. Next, we divide the number by ten (num = num / 10), which will remove the last digit from a number.

package main

import "fmt"

func getLength(num int) int {
	var count = 0
	for num > 0 {
		num = num / 10
		count += 1
	}

	return count
}

func main() {
	// var num, count int
	// count = 0
	// fmt.Scanf("%d ", &num)

	// for num > 0 {
	// 	num = num / 10
	// 	count += 1
	// }

	// fmt.Println("The total numbers are:", count)

	var num, count int
	fmt.Scanln(&num)
	count = getLength(num)
	fmt.Printf("The total Numbers are %d", count)

}
