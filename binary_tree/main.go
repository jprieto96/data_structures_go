package main

import "fmt"

type Node struct {
	value      int
	leftChild  *Node
	rightChild *Node
}

type BinaryTree struct {
	root *Node
}

func (binaryTree *BinaryTree) Insert(value int) *BinaryTree {
	if binaryTree.root == nil {
		binaryTree.root = &Node{value: value, leftChild: nil, rightChild: nil}
	} else {
		binaryTree.root.insert(value)
	}
	return binaryTree
}

func (node *Node) insert(value int) {
	if node == nil {
		return
	}
	if value <= node.value {
		if node.leftChild == nil {
			node.leftChild = &Node{value: value, leftChild: nil, rightChild: nil}
		} else {
			node.leftChild.insert(value)
		}
	} else {
		if node.rightChild == nil {
			node.rightChild = &Node{value: value, leftChild: nil, rightChild: nil}
		} else {
			node.rightChild.insert(value)
		}
	}
}

func main() {
	fmt.Println("HOlsaaaa")
}
