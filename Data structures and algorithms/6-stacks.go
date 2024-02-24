package main()

type Node struct {
	value int
	next *Node
}

type Stack struct {
	Top *Node
	Height int
}

func NewNode(value int) *Node {
	return &Node {
		value : value,
		next : nil,
	}
}

func NewStack() *Stack {
	return &Stack {
		Top : nil,
	}
}

func (s *Stack)GetTop() int {
	if s.Top != nil {
		return s.Top.value
	}
}

func (s *Stack)GetHeight() int {
	return s.Height
} 

func (s *Stack)Display() int {
	if s.Top == nil {
		return
	}
	current := s.Top

}

func (s *Stack)Push(value int) {
	newNode := NewNode(value)
	newNode.next = s.Top
	s.Top = newNode
	s.Height++
	
}

func (s *stack)Pop() int{
	
	if s.Height == 0 {
		return -1
	} 
	top := s.Top
	temp := top
	top = top.next
	temp.next = nil
	s.Height--
	return temp.value
}

func main(){

}