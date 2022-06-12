package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

//var used to call out of functions
var head *Node = new(Node)
var tail *Node = new(Node)

func linkedList() {
	head.value = 1
	head.next = nil

	var second *Node = &Node{
		value: 2,
		next:  nil,
	}
	head.next = second

	tail.value = 3
	tail.next = nil
	second.next = tail
}

func printList(node *Node) {
	for true {
		if node == nil {
			break
		}
		fmt.Printf("[%v]->", node.value)
		node = node.next
	}
}

func appendList(value int) {
	var newNode *Node = &Node{
		value: value,
		next:  nil,
	}

	tail.next = newNode
	tail = newNode
}

func prepend(value int) {
	var newNode *Node = &Node{
		value: value,
		next:  nil,
	}
	newNode.next = head
	head = newNode
}

func insert(pos int, value int) {
	i := 0
	start := head

	for true {
		if start.next == nil || i >= pos-1 {
			break
		}
		start = start.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.value = value
	newNode.next = start.next
	start.next = newNode
}

func removenode(pos int) {
	start := head
	i := 0

	for {
		if start.next == nil || i >= pos-1 {
			break
		}
		start = start.next
		i++
	}
	temp := start.next
	start.next = temp.next
	temp.next = nil
}

func main() {
	linkedList()
	appendList(4)
	prepend(0)
	insert(2, 6)
	removenode(3)
	printList(head)
}
