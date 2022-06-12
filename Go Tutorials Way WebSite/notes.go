/*Go Program to Count Total Notes in an Amount
Write a Go Program to Count the total number of
currency notes in a given amount using Arrays
and For loop. First, we declared an integer
array that holds the available notes.
Next, we used for loop (for i := 0; i < 8; i++)
to iterate the notes array and divides the
amount with each array item. Then, we update
the count by removing that cash amount from the
original.*/

package main

import "fmt"

func countNotes(amount int) {
	notes := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	temp := amount

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i], "Notes: ", temp/notes[i])
		temp = temp % notes[i]
	}
}

func main() {
	var amount int
	fmt.Println("Enter the Total Amount of Cash = ")
	fmt.Scanln(&amount)
	countNotes(amount)
}
