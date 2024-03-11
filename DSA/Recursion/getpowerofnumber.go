package main

import "fmt"

func PowerOf(x, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	if n%2 == 0 {
		return PowerOf(x*x, n/2)
	}
	return x * PowerOf(x, n-1)
}

func main() {
	fmt.Println(PowerOf(2, 4))
}
