package main

import "fmt"

type queue []int

func (q *queue) Push(data int) {
	*q = append(*q, data)
}
func (q *queue) Pop() {
	if len(*q) == 0 {
		fmt.Println("No Elements Found")
	} else {
		poppedEl := (*q)[0]
		*q = (*q)[1:]
		fmt.Println("Popped element : ", poppedEl)
	}
}
func (q *queue) Display() {
	fmt.Println(*q)
}
func main() {
	queue := &queue{}
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
}
