package main

import "fmt"

type LinkedList struct {
	Head *Node
}
type Node struct {
	Data int
	Next *Node
}

func (list *LinkedList) InsertData(data int) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
	} else {
		currentNode := list.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}
func (list *LinkedList) Display() {
	current := list.Head
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
func main() {
	list := &LinkedList{}
	list.InsertData(1)
	list.InsertData(2)
	list.InsertData(3)
	list.Display()
}
