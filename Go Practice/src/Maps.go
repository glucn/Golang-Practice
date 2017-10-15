package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count:=make(map[string]int)
	for _, word:= range strings.Fields(s) {
		count[word]=count[word]+1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
