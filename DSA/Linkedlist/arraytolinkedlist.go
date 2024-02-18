package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func ArrayToLinkedList(arr []int) *Node {
	var Head, Tail *Node
	for _, num := range arr {
		newNode := &Node{
			Data: num,
		}
		if Head == nil {
			Head = newNode
			Tail = newNode
		} else {
			Tail.Next = newNode
			Tail = newNode
		}
	}

	return Head
}
func Display(list *Node) {
	current := list
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
func main() {
	list := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	Display(list)
}
