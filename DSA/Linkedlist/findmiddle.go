package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
}

func (list *list) Push(data int) {
	newNode := &node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}
func (list *list) Display() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func (list *list) FindMiddle() *node {
	if list.head == nil {
		fmt.Println("List is empty")
	}
	node1 := list.head
	node2 := list.head
	for node2 != nil && node2.next != nil {
		node1 = node1.next
		node2 = node2.next.next
	}
	return node1
}
func main() {
	list := &list{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	// list.Push(6)

	list.Display()
	middle := list.FindMiddle()
	if middle != nil {
		fmt.Println("Middle element is : ", middle.data)
	}

}
