package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func NewNode(value int) *Node {
	return &Node{
		value : value,
		next : nil,
	}
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList)Display() {
	current := l.Head
	for current != nil {
		fmt.Println("%d->",current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList)Length() (length int) {
	current := l.Head
	//lenght := 0
	for current !=nil {
		length ++
		current = current.next
	}
	//return length 
	return
}

func (l *LinkedList)InsertNodeLastAppend(value int) {
	newNode := NewNode(value)
	//if list is empty
	if l.Head == nil {
		return
	}
	current := l.Head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList)InsertNodeFirstPrepend(value int) {
	newNode := NewNode(value)
	newNode.next = l.Head // add this 
	l.Head = newNode
}

func (l *LinkedList)DeleteNodeLast() {
	//if list is empty
	if l.Head == nil {
		return
	}
	//if list has only 1 element
	if l.Head.next == nil {
		l.Head = nil
	}
	current := l.Head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (l *LinkedList)DeleteNodeFirst(value int) {
	//if list is empty
	if l.Head == nil {
		return
	}
	//if list has only 1 element
	if l.Head.next == nil {
		l.Head = nil
	}

	current := l.Head
	if current != nil {
		current = current.next //deleting first element
	}
}

func (l *LinkedList)GetIndexValue(index int) int {
	length := l.Length()
	if index < 0 || index >= length {
		return 0
	}
	
	current := l.Head
	currentIndex := 0
	// for i:=0;i<index:i++ {
	// 	if i == index {
	// 		return current.value
	// 	}
	// }
	for current != nil {
		if currentIndex == index {
			return current.value
		}
		current = current.next
		currentIndex++
	}
	return current.value
}

func (l *LinkedList)SetIndexValue(index int, value int) {
	length := l.Length()
	if index < 0 || index >= length {
		return
	}

	current := l.Head
	currentIndex := 0

	// for  i:= 0; i<index; i++ {
	// 	if i == index {
	// 		current.value = value
	// 	}
	// }
	for current != nil {
		if currentIndex == index {
			current.value = value
		}
		current = current.next
		currentIndex++
	}
}

func (l *LinkedList)InsertNodeAtIndex(index int,value int) {
	length := l.Length()
	newNode := NewNode(value)
	current := l.Head
	currentIndex := 0
	if index < 0 || index >= length {
		return
	}

	//if insert at 0 index
	if index == 0 {
		newNode.next = current // point newNode's next to head
		current = newNode // point head to newnode
		return
	}

	for current != nil {
		if currentIndex == index-1 {
				newNode.next = current.next
				current.next = newNode
				return
		}
		current = current.next
		currentIndex++
	}
}

func (l *LinkedList)DeleteNodeAtIndex(index int,value int) {
	// invalid index
	if index < 0 { 
		return
	}
	//empty list
	if l.Head == nil { 
		return
	}
	//only 1 element
	if l.Head.next == nil {
		l.Head = nil
	}

	current := l.Head
	previous := current
	currentIndex:= 0
	for current != nil {
		if currentIndex == index { // index -1
			previous.next = current.next // current.next.next
		}
		previous = current
		current = current.next
		currentIndex++
	}
}

func (l *LinkedList)ReverseLinkedList() {
	//empty list
	if l.Head == nil { 
		return
	}
	//only 1 element
	if l.Head.next == nil {
		return 
	}

	var previous *Node
	current := l.Head
	var next *Node

	for current != nil {
		next = current.next
		current.next = previous
		previous = current 
		current = next
	}
	l.Head = previous // to point last node as head
}

func main () {

}