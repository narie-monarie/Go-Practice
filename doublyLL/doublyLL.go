package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func linkedList() {
	head.value = 1
	head.prev = nil
	head.next = nil

	var second *Node = &Node{
		value: 2,
		prev:  head,
		next:  nil,
	}
	head.next = second

	tail.value = 3
	tail.prev = second
	tail.next = nil
	second.next = tail
}

func printList(node *Node) {
	n := node
	var end *Node = nil
	for true {
		if n == nil {
			break
		}
		fmt.Printf("[%v]->", n.value)
		end = n
		n = n.next
	}
	fmt.Println("\n")

	n = end

	for true {
		if n == nil {
			break
		}
		fmt.Printf("[%v]->", n.value)
		n = n.prev
	}
}

func main() {
	linkedList()
	printList(head)
}
