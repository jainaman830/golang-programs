package main

import (
	"fmt"
	"time"
)

func worker(w int, jobs, results chan int) {
	for i := range jobs {
		time.Sleep(1 * time.Second)
		results <- i * 2 // Example processing: double the number
		fmt.Printf("Worker number : %d, Job Number : %d\n", w, i)
	}

}
func main() {
	fmt.Println("Program started")
	numOfJobs := 10
	numOfWorker := 3

	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for i := 0; i < numOfWorker; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < numOfJobs; i++ {
		jobs <- i + 1
	}
	close(jobs)
	for i := 0; i < numOfJobs; i++ {
		fmt.Println("result : ", <-results)
	}
}
