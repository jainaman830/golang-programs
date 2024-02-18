package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}
type Queue struct {
	head *node
	tail *node
}

func (q *Queue) PuchBack(data int) {
	newNode := &node{data: data}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		newNode.prev = q.tail
		q.tail = newNode
	}
}
func (q *Queue) Pop() {
	if q.head == nil {
		fmt.Println("no elements found")
	} else {
		poppedEl := q.head.data
		q.head = q.head.next
		fmt.Println("Popped element : ", poppedEl)
	}
}
func (q *Queue) Display() {
	current := q.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {
	queue := &Queue{}
	queue.PuchBack(1)
	queue.PuchBack(2)
	queue.PuchBack(3)
	queue.PuchBack(4)

	queue.Display()

	queue.Pop()
	queue.Display()
	queue.Pop()
	queue.Display()
}
