package main

import "strings"

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)
	word := strings.Split(trimmed, " ")
	if len(word) == 0 {
		return 0
	}
	return len(word[len(word)-1])
}
