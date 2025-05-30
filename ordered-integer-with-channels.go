package main

import (
	"fmt"
)

func printOdd(limit int, oddCh, evenCh chan int, done chan bool) {
	for {
		num, ok := <-oddCh
		if !ok {
			return
		}
		if num > limit {
			close(evenCh)
			done <- true
			return
		}
		fmt.Println(num)
		evenCh <- num + 1
	}
}

func printEven(limit int, evenCh, oddCh chan int, done chan bool) {
	for {
		num, ok := <-evenCh
		if !ok {
			return
		}
		if num > limit {
			close(oddCh)
			done <- true
			return
		}
		fmt.Println(num)
		oddCh <- num + 1
	}
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	oddCh := make(chan int)
	evenCh := make(chan int)
	done := make(chan bool, 2)

	go printOdd(n, oddCh, evenCh, done)
	go printEven(n, evenCh, oddCh, done)

	// Start with the first odd number
	oddCh <- 1

	// Wait for both goroutines to finish
	<-done
	<-done
}
