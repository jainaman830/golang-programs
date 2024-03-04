package main

import "fmt"

var (
	symbols = map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack []rune
)

func CheckExpression(s string) bool {
	stack = []rune{}
	if s == "" {
		return true
	}
	for _, c := range s {
		if _, ok := symbols[c]; ok {
			push(c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		el := pop()
		if symbols[el] != c {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
func push(c rune) {
	stack = append(stack, c)
}
func pop() rune {

	el := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return el
}
func main() {
	fmt.Println(CheckExpression("{{[]}}"))
	fmt.Println(CheckExpression("[{}}]"))
	fmt.Println(CheckExpression("({}[])"))
	fmt.Println(CheckExpression(")({[])"))
}
