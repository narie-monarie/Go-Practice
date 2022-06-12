package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

var root *Node = nil

func createNewNode(newValue int) *Node {
	var newNode *Node = new(Node)
	newNode.value = newValue
	newNode.left = nil
	newNode.right = nil
	return newNode
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("%d ", root.value)
	inOrder(root.right)
}

func add(node *Node, newValue int) {
	if root == nil {
		root = &Node{
			value: newValue,
			left:  nil,
			right: nil,
		}
		return
	}
	compareValue := newValue - node.value

	if compareValue < 0 {
		if node.left == nil {
			node.left = createNewNode(newValue)
		} else {
			add(node.left, newValue)
		}
	} else if compareValue > 0 {
		if node.right == nil {
			node.right = createNewNode(newValue)
		} else {
			add(node.right, newValue)
		}
	}
}

func main() {
	add(root, 60)
	add(root, 40)
	add(root, 20)
	add(root, 10)
	add(root, 30)
	add(root, 50)
	add(root, 80)
	add(root, 70)
	add(root, 90)
	fmt.Printf("in order: \n")
	inOrder(root)

}
