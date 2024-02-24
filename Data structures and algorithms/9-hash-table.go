package main

type Node struct {
	key string
	value int 
	next *Node
}

type HashTable struct {
	Bucket []*Node
	Size int
} 

func NewNode(key string, value int)*Node {
	return &Node {
		key : key,
		value : value,
		next : nil,
	}
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Bucket : make(Bucket []*Node, size), // remember
		Size : size,
	}
}

func  (h *HashTable)Hash(key string) int {
	hash := 0
	for _, char := range key {
			hash = hash + int(char)
	}
	index := hash % h.size
	return index

}

func (h *HashTable)Insert(key string, value int) {
	newNode := NewNode(key, value)
	index := Hash(key)

	if h.Bucket[index] == nil {
		h.Bucket[index] = newNode
	}
	newNode.next = h.bucket[index] // beginning of linked list always
	h.Bucket[index] = newNode
}

func(h *HashTable)Get(key string) int{
	index := Hash(key)
	temp := h.Bucket[index]
	for temp != nil {
		if temp.key == key {
			return temp.value
		}
		temp = temp.next
	}
	return 0
}

func (h *HashTable)Set(key string, value int) {
	newNode := NewNode(key, value)
	index := Hash(key)
	if h.Bucket[index] == nil {
		h.Bucket[index] = newNode
	} else {
		temp := h.Bucket[index]
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

func (h *HashTable)GetKeys()[]string {
	var key []string
	for _, node := range h.Bucket {
		for node != nil {
				key = append( keys, node.key)
				node = node.next
		}
	}
	return keys
}

func main() {
	
}