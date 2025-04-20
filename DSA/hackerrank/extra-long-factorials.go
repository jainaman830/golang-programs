/**/
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) {
	// Write your code here
	output := factorial(big.NewInt(int64(n)))
	fmt.Print(output)
}
func factorial(number *big.Int) (result *big.Int) {
	if number.Cmp(big.NewInt(1)) == 0 || number.Cmp(big.NewInt(1)) == -1 {
		return big.NewInt(1)
	}
	result = new(big.Int)
	result.Set(number)
	result.Mul(result, factorial(number.Sub(number, big.NewInt(1))))
	return
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
