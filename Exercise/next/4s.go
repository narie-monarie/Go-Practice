package main

import "fmt"

func enque(arr []int, element int) []int {
	arr = append(arr, element)
	return arr
}

func dequeue(arr []int) []int {
	element := arr[0]
	fmt.Println("We removed : ", element)
	return arr[1:]
}

func main() {
	arr := []int{4, 6, 7, 3, 4, 782, 24, 86, 2}
	arr = enque(arr, 3)
	arr = dequeue(arr)

	fmt.Println(arr)
}
