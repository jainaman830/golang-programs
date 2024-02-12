package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			// ch <- &value
			v := value
			ch <- &v
		}
		// close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // what will be printed here?
			wg.Done()
		}

	}()

	wg.Wait()
}
