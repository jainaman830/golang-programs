package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "slices"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	// slices.Sort(ranked)
	var currrank, max, min int32
	rankMap := make(map[int32]int32)
	for i := 0; i < len(ranked); i++ {
		if ranked[i] > max {
			max = ranked[i]
		}
		if ranked[i] < min {
			min = ranked[i]
		}
		if _, ok := rankMap[ranked[i]]; !ok {
			currrank++
			rankMap[ranked[i]] = currrank
		}
	}
	var arr []int32
	for _, num := range player {
		if rank, ok := rankMap[num]; ok {
			arr = append(arr, rank)
		} else {
			var lastRank int32
			if num > max {
				lastRank = 1
			} else if num < min {
				lastRank = int32(len(ranked) + 1)
			} else {
				for key, val := range rankMap {
					if num < key {
						if lastRank < val {
							lastRank = val
						}
					}
				}
				if lastRank == 0 {
					lastRank = 1
				} else {
					lastRank++
				}
			}

			arr = append(arr, lastRank)
			rankMap[num] = lastRank
		}
	}
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

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
