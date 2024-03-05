package main

import "fmt"

var stack []int

func LongestValidParentheses(s string) int {
	stack = []int{}
	push(-1)
	longest := 0
	for i, c := range s {
		if c == '(' {
			push(i)
			continue
		}
		//((()))
		pop()
		if len(stack) == 0 {
			push(i)
			continue
		}
		longest = max(longest, i-stack[len(stack)-1])
	}
	return longest
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func push(i int) {
	stack = append(stack, i)
}
func pop() {
	stack = stack[:len(stack)-1]
}
func main() {
	fmt.Println(LongestValidParentheses("(()))())("))
}
