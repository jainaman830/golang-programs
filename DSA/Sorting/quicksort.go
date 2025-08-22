//Concept: Picks a pivot and partitions the array into elements less than and greater than the pivot.
//Time Complexity: O(n log n)

package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
func main() {
	arr := []int{1, 39, 2, 9, 7, 54, 11, 1}
	fmt.Println(QuickSort(arr))
}
