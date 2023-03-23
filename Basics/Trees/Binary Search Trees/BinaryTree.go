package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) insert(val int) {
	newNode := &Node{
		value: val,
	}

	if b.root == nil {
		b.root = newNode
		return
	}

	currentNode := b.root

	for currentNode != nil {
		if val < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				return
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (b *BinaryTree) find(value int) bool {
	currentNode := b.root
	for currentNode != nil {
		if value == currentNode.value {
			return true
		} else if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return false
}

func (b *BinaryTree) BreadthFirstSearch() []int {
	node := b.root
	var data []int
	var queue []*Node
	queue = append(queue, b.root)

	for len(queue) != 0 {
		return data
	}
	data = append(data, queue[0].value)
}

func main() {
	myTree := BinaryTree{}
	myTree.insert(10)
	myTree.insert(6)
	myTree.insert(3)
	myTree.insert(8)
	myTree.insert(15)
	myTree.insert(20)

	fmt.Println(myTree.find(2))
	fmt.Println(myTree.find(20))
}
