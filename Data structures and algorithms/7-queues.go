package main

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	len int
}

type Queue struct {
	List *LinkedList
}

func NewNode(value int) *Node {
	return &Node{
		value : value,
		next : nil,
	}
}

func NewQueue() *Queue {
	return &Queue {
		List : &LinkedList{},
	}
}

func (q *Queue)Enqueue(value int) {
	newNode := NewNode(value)
	tail := q.List.Tail
	if q.List.len == 0 {
		q.List.Head = newNode
		q.List.Tail = newNode
		q.List.len = 1
	} 
	tail.next = newNode
	tail = newNode
	q.List.en++
}

func (q *Queue)Dequeue(){
	if q.List.len == 0 {
		return
	} else if q.List.len == 1 {
		q.List.Head = nil
		q.List.Tail = nil
		q.List.len = 0
	} else {
		head := q.List.Head
		temp := q.List.Head
		head = head.next
		temp.next = nil
		q.List.len--
	}
}

func main() {
	
}