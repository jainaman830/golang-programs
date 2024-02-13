// input = abcdefghijklmn
// key = 3
// output = cbafedihglkjnm
package main

import "fmt"

func ReverseSlice(arr []string) []string {
	for i := 0; i < (len(arr))/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
func main() {
	input := []string{"a", "b", "c", "d", "e", "f", "g"}
	k := 4
	reversArr := []string{}
	for i := 0; i < len(input); i += k {
		if k+i > len(input) {
			output := ReverseSlice(input[i:])
			reversArr = append(reversArr, output...)
		} else {
			output := ReverseSlice(input[i : k+i])
			reversArr = append(reversArr, output...)
		}

	}
	fmt.Println(reversArr)
}
