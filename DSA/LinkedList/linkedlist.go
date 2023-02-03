package main

type Node struct {
	value int
	next  *Node
}

type linkedlist struct {
	head   *Node
	length int
}

func (l *linkedlist) insert(val int) {
	//make a new node
	newNode := &Node{
		value: val,
	}

	if l.length == 0 {
		l.head = newNode
		l.length++
		return
	}

	current := l.head

	for i := 0; i < l.length; i++ {
		if current.next == nil {
			current.next = newNode
			l.length++
			return
		}
		current = current.next
	}
}
