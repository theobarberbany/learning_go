package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for i:= range words {
		val, prs := m[words[i]]
		if prs {
			m[words[i]] = val + 1
		} else {
			m[words[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

