package main

import "fmt"

type stack []int

func (s *stack) Insert(data int) {
	*s = append(*s, data)
}
func (s *stack) Pop() {
	if len(*s) == 0 {
		fmt.Println("No elements found")
	} else {
		poppedEl := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		fmt.Println("Popped element : ", poppedEl)
	}
}
func (s *stack) Display() {
	fmt.Println(*s)
}
func main() {
	stack := &stack{}
	stack.Insert(1)
	stack.Insert(2)
	stack.Display()
	stack.Pop()
	stack.Display()
}
