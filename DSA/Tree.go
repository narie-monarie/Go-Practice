package main

import "fmt"

type Node struct{
  value int
  left *Node
  right *Node
}

var root *Node = nil

func createNewNode(newValue int) *Node{
  var newNode *Node = new(Node)
  newNode.value = newValue
  newNode.left = nil
  newNode.right = nil
  return newNode
}

func inOrder(root *Node){
  if root == nil{
    return
  }
  inOrder(root.left)
  fmt.Printf("%d", root.value)
  inOrder(root.right)
}
