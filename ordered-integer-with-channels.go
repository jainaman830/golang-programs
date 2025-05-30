package main

import (
	"fmt"
	"sync"
)

func printOdd(limit int, oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range oddCh {
		if num > limit {
			close(evenCh)
			return
		}
		fmt.Println(num)
		evenCh <- num + 1
	}
}

func printEven(limit int, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range evenCh {
		if num > limit {
			close(oddCh)
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
	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(n, oddCh, evenCh, &wg)
	go printEven(n, evenCh, oddCh, &wg)

	oddCh <- 1 // Start with odd

	wg.Wait() // Wait for both goroutines to finish
}
