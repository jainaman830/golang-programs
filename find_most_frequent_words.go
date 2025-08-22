package main

import (
	"fmt"
	"strings"
)

// topKFrequent returns k most frequent words, breaking ties by first appearance
func topKFrequent(s string, k int) []string {
	words := strings.Fields(s)

	freq := make(map[string]int)      // word -> frequency
	firstSeen := make(map[string]int) // word -> first index of appearance

	for i, w := range words {
		freq[w]++
		if _, ok := firstSeen[w]; !ok {
			firstSeen[w] = i
		}
	}

	// Convert to slice for sorting
	type wordInfo struct {
		word      string
		frequency int
		index     int
	}

	wordList := []wordInfo{}
	for w, f := range freq {
		wordList = append(wordList, wordInfo{w, f, firstSeen[w]})
	}

	// Custom sort: higher frequency first, then earlier index
	// We’ll do a stable sort to keep tie-breaking predictable
	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			if wordList[i].frequency < wordList[j].frequency ||
				(wordList[i].frequency == wordList[j].frequency && wordList[i].index > wordList[j].index) {
				wordList[i], wordList[j] = wordList[j], wordList[i]
			}
		}
	}

	// Collect top k
	res := []string{}
	for i := 0; i < k && i < len(wordList); i++ {
		res = append(res, wordList[i].word)
	}

	return res
}

func main() {
	s := "i love go i love programming go go go"
	k := 3
	result := topKFrequent(s, k)
	fmt.Println(result) // Example output: ["go", "i", "love"]
}
