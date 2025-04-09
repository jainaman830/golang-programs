/*A video player plays a game in which the character competes in a hurdle race. Hurdles are of varying heights, and the characters have a maximum height they can jump. There is a magic potion they can take that will increase their maximum jump height by  unit for each dose. How many doses of the potion must the character take to be able to jump all of the hurdles. If the character can already clear all of the hurdles, return 0.*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	slices.Sort(height)
	diff := height[len(height)-1] - k
	if diff < 0 {
		return 0
	}
	return diff
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	heightTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var height []int32

	for i := 0; i < int(n); i++ {
		heightItemTemp, err := strconv.ParseInt(heightTemp[i], 10, 64)
		checkError(err)
		heightItem := int32(heightItemTemp)
		height = append(height, heightItem)
	}

	result := hurdleRace(k, height)

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
