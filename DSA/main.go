package main;

import "fmt"

type Node struct{
  value int
  left *Node
  right *Node

  Node (value int) {
    Node.value = value
    Node.left = Node.right = nil
  }
};

func preorderTraversal(node *Node){
  if (node == nil){
    return 
  }

  fmt.Println(node.value)
  preorderTraversal(node.left)
  preorderTraversal(node.right)
}
