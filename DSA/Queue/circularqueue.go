package main

import "fmt"

type Cqueue struct {
	Items    []interface{}
	capacity int
	start    int
	end      int
	size     int
}

func NewQueue(cap int) *Cqueue {
	queue := &Cqueue{
		Items:    make([]interface{}, cap),
		capacity: cap,
		start:    0,
		end:      -1,
		size:     0,
	}
	return queue
}
func (q *Cqueue) Push(data int) {
	if q.size == q.capacity {
		fmt.Println("Queue is full")
		return
	}
	q.end = (q.end + 1) % q.capacity
	q.Items[q.end] = data
	q.size++
}
func (q *Cqueue) Pop() {
	if q.size == 0 {
		fmt.Println("Queue is empty")
		return
	}
	elem := q.Items[q.start]
	q.Items[q.start] = nil
	q.start = (q.start + 1) % q.capacity
	q.size--
	fmt.Println("Popped element : ", elem)
}
func (q *Cqueue) Display() {
	fmt.Println(q.Items)
}
func main() {
	queue := NewQueue(5)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)

	queue.Display()
	queue.Pop()
	queue.Display()
	queue.Pop()
	queue.Display()
	queue.Push(6)
	queue.Display()
}
