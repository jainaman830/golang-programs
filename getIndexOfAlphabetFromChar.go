package main

import "fmt"

func main() {
	numeral := alphaNum('T')
	fmt.Println(numeral)
}

func alphaNum(letter rune) int {
	char := int(letter)
	char -= 64
	return char
}
package main

import "fmt"

func main() {
	numeral := alphaNum('T')
	fmt.Println(numeral)
}

func alphaNum(letter rune) int {
	char := int(letter)
	char -= 64
	return char
}
