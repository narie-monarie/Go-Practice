package main

import "fmt"

func main() {
	var array1 [5]int
	array := [...]int{1, 2, 4, 1, 4}
	array[3] = 35

	//Initializing the array position
	pos := [6]int{1: 20, 3: 40}
	fmt.Println(pos)
	array1 = array
	fmt.Println(array1)
	fmt.Println(array[3])

	var arr1 [3]*string
	arr := [3]*string{new(string), new(string), new(string)}

	*arr[0] = "Red"
	*arr[1] = "Blue"
	*arr[2] = "Green"

	arr1 = arr

	fmt.Println(arr1)

	//multidimension arrays

	arr2 := [2][4]int{{10, 11, 12, 13}, {14, 15, 16, 17}}
	arr2[0][1] = 21
	fmt.Println(arr2)

}
