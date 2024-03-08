package main

import "fmt"

func ReverseNumber(num int) int {
	output := 0

	for num != 0 {
		digit := num % 10
		output = output*10 + digit
		num = num / 10
	}
	return output
}

func main() {
	fmt.Println(ReverseNumber(-11002))
}
