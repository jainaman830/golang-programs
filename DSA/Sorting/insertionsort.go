// Concept: Builds the sorted array one item at a time, like sorting playing cards.
// Time Complexity : O(n2)
package main

import "fmt"

func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		el := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > el {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}
	return arr
}
func main() {
	arr := []int{1, 39, 2, 9, 7, 54, 11, 1}
	/*
	   [1 39 2 9 7 54 11 1]
	   [1 2 39 9 7 54 11 1]
	   [1 2 9 39 7 54 11 1]
	   [1 2 7 9 39 54 11 1]
	   [1 2 7 9 39 54 11 1]
	   [1 2 7 9 11 39 54 1]
	   [1 1 2 7 9 11 39 54]
	   [1 1 2 7 9 11 39 54]
	*/
	fmt.Println(insertionSort(arr))
}
