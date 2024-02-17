package main

import "fmt"

func PrintEven(even chan int) {
	for {
		fmt.Println(<-even)
	}
}
func CheckEven(input, even chan int, done chan bool) {
	for num := range input {
		num = <-input
		if num%2 == 0 {
			even <- num
		}
	}
	done <- true
}
func main() {
	n := 10
	input := make(chan int, n)
	for i := 1; i <= n; i++ {
		input <- i
	}
	close(input)
	done := make(chan bool)
	even := make(chan int)
	go CheckEven(input, even, done)
	go PrintEven(even)
	<-done
}
