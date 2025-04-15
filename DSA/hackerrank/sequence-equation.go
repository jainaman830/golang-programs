/*
Given a sequence of n integers, p(1), p(2),...,p(n) where each element is distinct and satisfies 1 ≤ p(x) ≤ n. For each x where 1 ≤ x ≤ n, that is x increments from 1 to n, find any integer y such that p(p(y)) ≡ x and keep a history of the values of y in a return array.
p = [5, 2, 1, 3, 4]
Each value of x between 1 and 5, the length of the sequence, is analyzed as follows:
Function Description
Complete the permutationEquation function in the editor below.

permutationEquation has the following parameter(s):

int p[n]: an array of integers
Returns
int[n]: the values of y for all x in the arithmetic sequence 1 to n
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var arr []int32
	numMap := make(map[int32]int32)
	for i, num := range p {
		numMap[num] = int32(i + 1)
	}
	for i := 0; i < len(p); i++ {
		index := numMap[int32(i+1)]
		index2 := numMap[index]
		arr = append(arr, index2)
	}
	return arr
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

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := permutationEquation(p)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
