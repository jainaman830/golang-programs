package main

import (
	"fmt"
	"strings"
)

func CheckPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return false
	} else if len(s) == 1 {
		return true
	} else if s[0] != s[len(s)-1] {
		return false
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func CheckPalindromeUsingRecursion(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return false
	} else if len(s) == 1 {
		return true
	} else if s[0] != s[len(s)-1] {
		return false
	}
	return CheckPalindromeUsingRecursion(s[1 : len(s)-1])
}
func main() {
	fmt.Println(CheckPalindrome("NAN1N"))
	fmt.Println(CheckPalindromeUsingRecursion("NAN"))
}
