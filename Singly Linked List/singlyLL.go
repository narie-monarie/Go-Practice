package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

var head *Node = new(Node) //creating the head Node
var tail *Node = new(Node)

func initial() {
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

	var tail *Node = &Node{
		value: 4,
		next:  nil,
	}
	third.next = tail
}

func appendval(value int) {
	n := head
	for {
		if n == nil {
			break
		}
		n = n.next
	}
	var newNode *Node = new(Node)
	n = newNode
	newNode.next = n
}

func output(node *Node) {
	p := node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%d->", p.value)
		p = p.next
	}
}

func main() {
	initial()
	appendval(2)
	output(head)
}
