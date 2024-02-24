package main

type Node struct {
	value int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	len int
}

func NewNode(value int) *Node {
	return &Node{
		value : value,
		next : nil,
		prev : nil
	}
}

func (d *DoublyLinkedList)InsertFirstPrepend(value int) {
	newNode := NewNode(value)
	if d.len == 0 {
		d.Head = newNode
		d.Tail = newNode
		d.len = 1
	} else {
		head := d.Head
		head.prev = newNode
		newNode.next = head
		head = newNode
		d.len++
	}
}

func (d *DoublyLinkedList)InsertEndAppend(value int) {
	newNode := NewNode(value)
	if d.len == 0 {
		d.Head = newNode
		d.Tail = newNode
		d.len = 1
	} else {
		tail = d.Tail
		newNode.prev = tail
		tail.next = newNode
		tail = newNode
		d.len++
	}
}

func (d *DoublyLinkedList)InserNodeAtIndex(index int, value int) {
	newNode := NewNode(value)
	if index < 0 || index >= d.len {
		return
	}
	var current *Node
	if index < d.len / 2 {
		current = l.Head
		currentIndex := 0
		for current != nil {
			if currentIndex == index - 1 {
				temp := current.next
				current.next = newNode
				newNode.next = temp
				temp.prev = newNode
				newNode.prev = current
				d.len++
			}
			current = current.next
			currentIndex++
		}
	} else  {
		current = l.Tail
		currentIndex := d.len - 1
		for current ! = nil {
			if currentIndex == index + 1 {
				temp := current.prev
				temp.next = newNode
				current.prev = newNode
				newNode.prev = temp
				newNode.next = current
				d.len++
			}
			current = current.prev
			currentIndex-- 
		}
	}
}

func (d *DoublyLinkedList)DeleteFirstNode(){
	if d.len == 0 {
		return
	} 
	if d.len == 1 {
		d.Head = nil
		d.Tail = nil
		d.len = 0
	}
	head := d.Head
	head  = head.next
	head.prev.next = nil
	head.prev = nil
	d.len--
}

func (d *DoublyLinkedList)DeleteLastNode() {
	if d.len == 0 {
		return
	}
	if d.len == 1 {
		d.Head = nil
		d.Tail = nil
		d.len = 0
	}
	tail = d.Tail
	tail = tail.prev
	tail.next.prev = nil
	tail.next = nil
	d.len--
}

func (d *DoublyLinkedList)DeleteNodeAtIndex(index int) {
	if d.len == 0 || index < 0 || index >= 0 {
		return
	}
	var current *Node
	if index < d.len/2 {
		current = d.Head
		currentIndex := 0
		for current != nil {
			if currentIndex == index-1 {
				temp := current.next
				current.next = temp.next
				temp.next.prev = current
				temp.next = nil
				temp.prev = nil
				d.len--
			}
			current = current.next
			currentIndex++
		}
	} else {
		current := d.Tail
		currentIndex := d.len-1
		for current != nil {
			if currentIndex  == index + 1 {
				temp := current.prev
				current.prev = temp.prev
				tem.prev.next = current
				temp.prev = nil
				temp.next = nil
				d.len--
			}
			current = current.prev
			currentIndex--
		}
	}
}

func (d *DoublyLinkedList)GetIndexValue(index int) int {
	if index < 0 || index >= d.len {
		return
	}
	var current *Node
	if index < d.len/2 {
		current = d.Head
		currentIndex := 0
		for current != nil {
			if currentIndex == index {
				return current.value
			}
			current = current.next
			currentIndex++
		}
	} else {
		current = d.Tail
		currentIndex := d.len - 1
		for current != nil {
			if currentIndex == index {
				return current.value
			}
			current = current.prev
			currentIndex--
		}
	}
}

func (d *DoublyLinkedList)SetIndexValue(index int, value int) {
	if index < 0 || index >= d.len {
		return
	}
	var current *Node
	if index < d.len/2 {
		current = d.Head
		currentIndex := 0
		for current != nil {
			if currentIndex == index {
				current.value = value
			}
			current = current.next
			currentIndex++
		}
	} else {
		current = d.Tail
		currentIndex := d.len - 1
		for current != nil {
			if currentIndex == index {
				current.value = value
			}
			current = current.prev
			currentIndex--
		}
	}
}

func (d *DoublyLinkedList)Display() {
	if d.len == 0 {
		fmt.Println("Empty")
	} else {
		fmt.Println("Forward: ")
		for current != nil {
			fmt.Print("%d->",current.value)
			current = current.next
		}
		fmt.Print("nil")

		fmt.Println("Backward: ")
		for current != nil {
			fmt.Print("%d->",current.value)
			current = current.prev
		}
		fmt.Print("nil")
	}
} 

func (d *DoublyLinkedList)ReverseDoublyLinkedList() {
	if d.len == 0 || d.len == 1 {
		return
	}
	current := d.Head
	d.Head = d.Tail
	d.Tail = d.Head
	for current != nil {
		current.next = current.prev
		current.prev = current.next
		current = current.prev
	}
}

func (d *DoublyLinkedList)Length()(length int){
	//return d.len
	//if len is not defined 
	current := d.Head
	for current != nil {
		length ++
		current = current.next
	}
	return 
}

func main() {
	
}