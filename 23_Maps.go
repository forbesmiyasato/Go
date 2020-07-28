package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordMap := make(map[string]int)
	for index, value := range words {
		wordMap[value]++
		_ = index
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
