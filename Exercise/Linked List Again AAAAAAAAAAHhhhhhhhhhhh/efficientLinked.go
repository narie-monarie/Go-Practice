package main

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedlist) First() *node {
	return l.head
}

func (l *linkedlist) prepend(val int) {
	newNode := node{val: val}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func (l *linkedlist) Push(val int) {
	newNode := &node{val: val}
	if l.head == nil {
		l.head = newNode
		l.length++
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.length++

}
func (l *linkedlist) printList() {
	n := l.head
	for n != nil {
		fmt.Printf("%d->", n.val)
		n = n.next
	}
}

func (l *linkedlist) deleteWithVal(val int) {
	if l.length == 0 {
		return
	}

	if l.head.val == val {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head

	for prevToDelete.next.val != val {
		if prevToDelete.next.next == nil {
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func main() {
	myList := linkedlist{}
	myList.Push(2)
	myList.prepend(12)
	myList.prepend(11)
	myList.prepend(7)
	myList.prepend(55)
	myList.deleteWithVal(12)
	myList.printList()
}
