package main

import "fmt"

func ReverseNumber(num, output int) int {
	if num != 0 {
		digit := num % 10
		output = output*10 + digit
		num = num / 10
		return ReverseNumber(num, output)
	}
	return output

}

func main() {
	output := ReverseNumber(-11002, 0)
	fmt.Println(output)
}
