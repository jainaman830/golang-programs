package main

import "fmt"

func selectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		fmt.Println(min, arr)
	}
	return arr
}
func main() {
	arr := []int{1, 39, 2, 9, 7, 54, 11, 1}
	fmt.Println(selectionSort(arr))
}
