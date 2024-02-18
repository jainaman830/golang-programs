package main

import "fmt"

type node struct {
	data int
	next *node
}

type StackLinkedList struct {
	head *node
	size int
}

func (s *StackLinkedList) Insert(data int) {
	s.head = &node{data: data, next: s.head}
	s.size++
}
func (s *StackLinkedList) Pop() {
	poppedEl := s.head.data
	fmt.Println("Popped element : ", poppedEl)
	s.head = s.head.next
	s.size--
}
func (s *StackLinkedList) Display() {
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {
	stack := &StackLinkedList{}
	stack.Insert(1)
	stack.Insert(2)
	stack.Insert(3)
	stack.Insert(4)
	stack.Display()
	stack.Pop()
	stack.Display()
}
