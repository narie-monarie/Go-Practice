package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	intList.PushFront(3)
	fmt.Println(intList.Back().Next())
	fmt.Println(intList.Len())
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d->", element.Value.(int))

	}

}
