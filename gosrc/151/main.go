package main

import (
	"fmt"
	"strings"
)

func reverse(words []string) []string {
	var l = len(words)
	for i := 0; i < l-i-1; i++ {
		var tmp = words[i]
		words[i] = words[l-i-1]
		words[l-i-1] = tmp
	}
	return words
}

func reverseWords(s string) string {
	var words []string
	var prev = 0
	for i, c := range s {
		if c == ' ' {
			if i > prev {
				words = append(words, s[prev:i])
			}
			prev = i + 1
		}
	}
	if len(s) > prev {
		words = append(words, s[prev:len(s)])
	}
	return strings.Join(reverse(words), " ")
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}
