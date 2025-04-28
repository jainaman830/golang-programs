package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	// Write your code here
	s = strings.ReplaceAll(s, " ", "")
	sqrt := math.Sqrt(float64(len(s)))
	rows := math.Floor(sqrt)
	cols := math.Ceil(sqrt)
	fmt.Println(sqrt, rows, cols)
	arr := []string{}
	i := 0
	for i < len(s) {
		end := i + int(cols)
		if end > len(s) {
			end = len(s)
		}
		arr = append(arr, s[i:end])
		i = end
	}
	output := ""
	for i := 0; i < int(cols); i++ {
		if output != "" {
			output += " "
		}
		for _, str := range arr {
			if i < len(str) {
				fmt.Println(str, i)
				output += string(str[i])
			}
		}
	}
	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

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
