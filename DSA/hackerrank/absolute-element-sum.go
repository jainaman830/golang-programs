package main

import (
	"fmt"
	"sort"
)

func summation(arr []int, queries []int) {
	l := len(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	cumulative := make([]int, l+1)
	for i := 0; i < l; i++ {
		cumulative[i+1] = cumulative[i] + arr[i]
	}
	fmt.Println(cumulative)
	x := 0
	for _, el := range queries {
		x += el
		i := sort.Search(l, func(i int) bool { return arr[i] >= -x })
		n_bigger := l - i
		n_smaller := i
		fmt.Println("x:", x, "i:", i, "n_bigger:", n_bigger, "n_smaller:", n_smaller)
		sum1 := n_bigger*x + cumulative[l] - cumulative[i]
		sum2 := -cumulative[i] - n_smaller*x
		fmt.Println(sum1 + sum2)
	}
}
func main() {
	arr := []int{-1, 2, -3}
	queries := []int{1, -2, 3}
	summation(arr, queries)
}
