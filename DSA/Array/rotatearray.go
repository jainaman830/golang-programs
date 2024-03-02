package main

import "fmt"

func rotateRight(arr []int, count int) []int {
	length := len(arr)
	for c := 0; c < count; c++ {
		last := arr[length-1]
		for i := length - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = last
	}

	return arr
}

func rotateLeft(arr []int, count int) []int {
	length := len(arr)
	for c := 0; c < count; c++ {
		first := arr[0]
		for i := 0; i < length-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[length-1] = first
	}
	return arr
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(rotateRight(arr, 7))
	fmt.Println(rotateLeft(arr, 7))
}
