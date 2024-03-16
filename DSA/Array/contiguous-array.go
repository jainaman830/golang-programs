package main

import "fmt"

func findMaxLength(nums []int) int {
	var count, max int
	var hash = map[int]int{0: -1}
	for i, num := range nums {
		if num == 0 {
			count -= 1
		} else {
			count += 1
		}
		fmt.Print(i, " num : ", num, ", count : ", count, " before : ", hash)
		if index, ok := hash[count]; ok {
			max = getMax(max, i-index)
		} else {
			hash[count] = i
		}
		fmt.Print(" After : ", hash, ", max : ", max, "\n")
	}
	return max
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 1, 0}))
}
