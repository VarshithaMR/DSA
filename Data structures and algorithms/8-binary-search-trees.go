package main

type Node struct {
	value int
	Left *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewNode(value int) *Node {
	return &Node{
		value : value,
	}
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree {}
}

func (b *BinarySearchTree)Insert(value int) {
	newnode := NewNode(value)
	if b.Root == nil {
		b.Root = newNode
	}
	temp := b.Root
	for {
		if newNode.value == temp.value {
			return
		} else if newNode.value < temp.value {
			if temp.Left == nil {
				temp.Left = newNode
				return
			}
			temp = temp.Left
		} else {
			if temp.Right == nil {
				temp.Right = newNode
				return
			}
			temp = temp.right
		}
	}	
}

func (b *BinarySearchTree)Contains(value int) bool {
	temp := b.root
	for temp != nil {
		if value < temp.value {
			temp = temp.Left
		} else if value > temp.value {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
} 

func main() {

}