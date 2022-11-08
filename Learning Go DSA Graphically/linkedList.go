package main

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) First() *Node {
	return l.head
}

func (l *LinkedList) AddFront(value int) {
	node := &Node{value: value}
  if l.head == nil{
    l.head = node
  }else{
    node.next = l.head
    l.head = node
  }
  return

}

func (l *LinkedList) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *LinkedList) Pop() {
	if l.head == nil {
		return
	} else {
		l.tail = nil
	}
}

func main() {
	l := &LinkedList{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.AddFront(7)

	n := l.First()

	for {
		println(n.value)
		n = n.next
		if n == nil {
			break
		}
	}
}
