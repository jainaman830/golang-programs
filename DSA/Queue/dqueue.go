package main

import "fmt"

type node struct {
	value    int
	priority int
	next     *node
}
type queue struct {
	head *node
}

func (q *queue) Push(data, priority int) {
	newNode := &node{value: data, priority: priority}
	if q.head == nil || priority < q.head.priority {
		newNode.next = q.head
		q.head = newNode
	} else {
		current := q.head
		for current.next != nil && priority >= current.next.priority {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
}
func (q *queue) Pop() {
	if q.head == nil {
		return
	}
	poppedEl := q.head.value
	q.head = q.head.next
	fmt.Println("Popped element : ", poppedEl)
}
func (q *queue) Display() {
	current := q.head
	for current != nil {
		fmt.Println(current.value, current.priority)
		current = current.next
	}
}
func main() {
	q := &queue{}
	q.Push(1, 1)
	q.Push(2, 2)
	q.Push(3, 3)
	q.Push(4, 1)
	q.Push(5, 1)
	q.Push(6, 2)

	q.Display()

	q.Pop()
	q.Display()
}
