package main

import "fmt"

type any interface {
}

type Node struct {
	value      *any
	leftChild  *Node
	rightChild *Node
}

type BinaryTree struct {
	root *Node
}

func (binaryTree *BinaryTree) insert(value *any) {
	if binaryTree != nil {

	}
}

func main() {
	fmt.Println("HOlsaaaa")
}
