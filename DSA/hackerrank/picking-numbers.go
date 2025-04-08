/*
Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to 1.

Example
a=[1,1,2,2,4,4,5,5,5]

There are two subarrays meeting the criterion: [1,1,2,3] and . [4,4,5,5,5]The maximum length subarray has  5 elements.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
* Complete the 'pickingNumbers' function below.
*
* The function is expected to return an INTEGER.
* The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here
	slices.Sort(a)
	maxLen := 0
	for i, num := range a {
		var arr []int32
		arr = append(arr, num)
		smallestNum := num
		for j, num2 := range a {
			if i != j {
				if math.Abs(float64(num-num2)) <= 1 && math.Abs(float64(smallestNum-num2)) <= 1 {
					arr = append(arr, num2)
					if smallestNum > num2 {
						smallestNum = num2
					}
				}

			}
		}
		if len(arr) > maxLen {
			maxLen = len(arr)
		}
		if len(arr) == len(a) {
			break
		}
	}
	return int32(maxLen)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
