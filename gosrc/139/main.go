package main

import (
	"fmt"
	"sort"
)

type Words struct {
	sort.Interface
	wordDict []string
}

func (ws Words) Len() int {
	return len(ws.wordDict)
}

func (ws Words) Swap(i, j int) {
	var tmp = ws.wordDict[i]
	ws.wordDict[i] = ws.wordDict[j]
	ws.wordDict[j] = tmp
}

func (ws Words) Less(i, j int) bool {
	return (len(ws.wordDict[i]) < len(ws.wordDict[j]))
}

func (ws Words) String() string {
	return fmt.Sprintln(ws.wordDict)
}

var mem []bool

func dfs(s string, wordDict Words) bool {
	var ls = len(s)
	if ls == 0 {
		return true
	}
	if mem[ls] {
		return false
	}

	for _, word := range wordDict.wordDict {
		var lw = len(word)
		if ls >= lw {
			if s[:lw] == word {
				var b = dfs(s[lw:], wordDict)
				if b {
					return true
				} else {
					mem[ls] = true
				}
			}
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	mem = make([]bool, len(s)+1)
	var words = Words{wordDict: wordDict}
	sort.Sort(words)
	return dfs(s, words)
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}
