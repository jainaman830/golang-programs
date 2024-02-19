package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		lastNode := ll.head
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = newNode
	}
}

func (list *LinkedList) removeEven() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}
	for list.head != nil && list.head.data%2 == 0 {
		list.head = list.head.next
	}
	current := list.head
	for current != nil && current.next != nil {
		if current.next.data%2 == 0 {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	ll.insert(12)
	ll.insert(2)
	ll.insert(3)
	ll.insert(4)
	ll.insert(6)
	ll.insert(8)
	ll.insert(9)

	fmt.Println("Original Linked List:")
	ll.display()

	ll.removeEven()
	fmt.Println("Linked List after removing even numbers:")
	ll.display()
}
