package main

import (
	"fmt"
	"math"
)

func minOperations(price []int, query []int) []int {
	result := make([]int, len(query))

	for i, q := range query {
		operations := 0
		for _, p := range price {
			operations += int(math.Abs(float64(p - q)))
		}
		result[i] = operations
	}

	return result
}

func main() {
	//a shop in hackerland contains n items where the price of ith items is price[i]. in one operation the price of any one item can be increased or decreased by 1.given q queries denoted by the array query find the minimum number of operations required to make the price of all items equal query[i]
	price := []int{1, 2, 3}
	query := []int{3, 2, 1, 5}
	result := minOperations(price, query)
	fmt.Println("Minimum number of operations required for each query:", result)
}
