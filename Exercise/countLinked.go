package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initialize() {
	head.value = 1
	head.next = nil

	var second *Node = &Node{
		value: 2,
		next:  nil,
	}
	head.next = second

	var third *Node = &Node{
		value: 3,
		next:  nil,
	}
	second.next = third

	tail.value = 4
	tail.next = nil
	third.next = tail
}

func printList(node *Node) {
	for true {
		if node == nil {
			break
		}
		fmt.Printf("[%d]->", node.value)
		node = node.next
	}
}

func remove(node *Node) {
	if node == nil {
		return
	}

	temp := node

	for temp.next.next != nil {
		temp.value = temp.next.value
		temp = temp.next
	}
	temp.value = temp.next.value
	temp.next = nil
}
func main() {
	initialize()
	printList(head)
	remove(head.value)
}
