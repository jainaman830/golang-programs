package main

import "fmt"

type LinkedList struct {
	Head *Node
}
type Node struct {
	Data int
	Next *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}
func (list *LinkedList) Reverse() {
	var prev *Node = nil
	current := list.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}
func (list *LinkedList) Printlist() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
func main() {
	list := &LinkedList{}
	for i := 1; i <= 5; i++ {
		list.Insert(i)
	}
	list.Printlist()
	list.Reverse()
	fmt.Println("After reverse")
	list.Printlist()
}
