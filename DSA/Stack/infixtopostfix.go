package main

import "fmt"

var operands = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
}

type operators struct {
	stack []string
}

func InfixtoPrefix(infix []string) []string {
	output := []string{}
	stack := new(operators)
	for _, c := range infix {
		if _, ok := operands[c]; !ok && c != "(" && c != ")" {
			output = append(output, c)
			continue
		} else if c == "(" {
			stack.push(c)
		} else if c == ")" {
			for len(stack.stack) > 0 && stack.stack[len(stack.stack)-1] != "(" {
				el := stack.pop()
				output = append(output, el)
			}
			stack.pop() //discard "("
		} else {
			for len(stack.stack) > 0 && operands[c] <= operands[stack.stack[len(stack.stack)-1]] {
				el := stack.pop()
				output = append(output, el)
			}
			stack.push(c)
		}
	}
	for len(stack.stack) > 0 {
		el := stack.pop()
		output = append(output, el)
	}
	return output
}
func (operator *operators) push(c string) {
	operator.stack = append(operator.stack, c)
}
func (operator *operators) pop() string {
	el := operator.stack[len(operator.stack)-1]
	operator.stack = operator.stack[:len(operator.stack)-1]
	return el
}
func main() {
	infixExpr := []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2"}
	postfixExpr := InfixtoPrefix(infixExpr)
	fmt.Println("Infix Expression:", infixExpr)
	fmt.Println("Postfix Expression:", postfixExpr)
}
