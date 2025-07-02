//Efficient search on sorted arrays by dividing the search space in half
//time complexity : O(logn)

package main

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == val {
			return mid
		} else if val < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1 // Value not found
}
func main() {

}
