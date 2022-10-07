package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

var head *Node = new(Node)

func initialize() {
	head.value = 7
	head.next = nil
}
func append(value int) {
	var tail *Node = &Node{
		value: value,
		next:  nil,
	}
	head.next = tail
}

func prepend(value int) {
	var newHead *Node = &Node{
		value: value,
		next:  nil,
	}
	newHead.next = head
	head = newHead
}
func insert_Position(pos int, value int) {
	n := head
	i := 0

	for {
		if n.next == nil || i >= pos-1 {
			break
		}
		n = n.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.value = value
	newNode.next = n.next
	n.next = newNode

}
func printList(n *Node) {
	for {
		if n != nil {
			fmt.Printf("%d->", n.value)
			n = n.next
		}
	}
}
func main() {
	initialize()
	append(12)
	insert_Position(1, 23)
	prepend(1)
	printList(head)
}
