package main

import (
	"fmt"
	"strconv"
)

func EvaluatePostfix(str string) int {
	stack := []int{}
	for _, c := range str {
		token := string(c)

		if token == "+" || token == "-" || token == "/" || token == "*" {
			if len(stack) < 2 {
				return -1
			}
			val2 := stack[len(stack)-1]
			val1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			output := 0
			switch token {
			case "+":
				output = val1 + val2
			case "-":
				output = val1 - val2
			case "*":
				output = val1 * val2
			case "/":
				output = val1 / val2
			}
			stack = append(stack, output)
		} else {
			val, err := strconv.Atoi(token)
			if err != nil {
				return -1
			}
			stack = append(stack, val)
		}
	}
	if len(stack) == 1 {
		return stack[0]
	} else {
		return -1
	}
}

func main() {
	expression := "543*+2-"
	result := EvaluatePostfix(expression)
	fmt.Println("Result:", result) // Output: Result: 15
}
