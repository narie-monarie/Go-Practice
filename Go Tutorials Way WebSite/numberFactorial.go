//calculate the factorial of a number

package main

import "fmt"

func main() {
	num := 1
	ranger := 5

	for i := 1; i <= ranger; i++ {
		num *= i
	}
	fmt.Println(num)
}
