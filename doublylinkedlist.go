package main

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
}
type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func (list *LinkedList) InsertAtLast(data int) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
}
func (list *LinkedList) InsertAtFront(data int) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
}

func (list *LinkedList) Display() {
	currentNode := list.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}
func (list *LinkedList) DisplayReverse() {
	fmt.Println("Print reverse")
	currentNode := list.Tail
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Prev
	}
}
func (list *LinkedList) RemoveFirst() {
	data := list.Head.Data
	next := list.Head.Next
	next.Prev = nil
	list.Head = next
	fmt.Println("Removed : ", data)
}
func (list *LinkedList) RemoveLast() {
	data := list.Tail.Data
	prev := list.Tail.Prev
	prev.Next = nil
	list.Tail = prev
	fmt.Println("Removed : ", data)
}
func main() {
	list := &LinkedList{}
	list.InsertAtFront(1)
	list.InsertAtFront(2)
	list.InsertAtLast(3)
	list.InsertAtFront(4)

	list.Display()
	list.DisplayReverse()

	list.RemoveFirst()
	list.Display()
	list.RemoveLast()
	list.Display()
}
