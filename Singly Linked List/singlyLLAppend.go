package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

var head *Node = new(Node) //creating the head Node
var tail *Node = new(Node) //creating the tail Node

func initial() {
	head.value = 10
	head.next = nil

	var second *Node = &Node{value: 12, next: nil}
	head.next = second

	var third *Node = &Node{value: 31, next: nil}
	second.next = third

	tail.value = 43
	tail.next = nil
	third.next = tail
}

func addEnd(value int) {
	var newNode *Node = &Node{
		value: value,
		next:  nil,
	}
	tail.next = newNode
	tail = newNode
}

func addStart(value int) {
	var newNode *Node = &Node{
		value: value,
		next:  nil,
	}
	newNode.next = head
	head = newNode

}

func searchIndex(position int) int {
	index := 0
	temp := head
	for temp.value != position {
		index++
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("Not Found")
	}

	return index
}

func insert(position, value int) {
	temp := head
	index := 0

	for {
		if temp.next == nil || index >= position-1 {
			break
		}
		temp = temp.next
		index++
	}
	var newNode *Node = new(Node)
	newNode.value = value
	newNode.next = temp.next
	temp.next = newNode
}
func output(node *Node) {
	for {
		if node == nil {
			break
		}
		fmt.Printf("[%d]->", node.value)
		node = node.next
	}
}
func main() {
	appender := 5
	initial()
	addEnd(appender)
	addStart(appender)
	fmt.Println(searchIndex(12))
	insert(3, 27)
	output(head)
}
