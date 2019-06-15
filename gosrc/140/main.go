package main

import (
	"fmt"
	"sort"
	"strings"
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

var mem map[int][][]string
var impossible []bool

func dfs(s string, wordDict Words, curr []string) [][]string {
	var ls = len(s)
	if impossible[ls] {
		return nil
	}
	if data, ok := mem[ls]; ok {
		return data
	}

	var res [][]string
	for _, word := range wordDict.wordDict {
		var lw = len(word)
		if ls >= lw {
			if s[:lw] == word {
				if ls == lw {
					res = append(res, []string{word})
				} else {
					var newCurr = append(curr, word)
					for _, suffix := range dfs(s[lw:], wordDict, newCurr) {
						res = append(res, append([]string{word}, suffix...))
					}
				}
			}
		} else {
			break
		}
	}

	if res == nil {
		impossible[ls] = true
	} else {
		mem[ls] = res
	}
	return res
}

func wordBreak(s string, wordDict []string) []string {
	impossible = make([]bool, len(s)+1)
	mem = make(map[int][][]string)
	var words = Words{wordDict: wordDict}
	sort.Sort(words)
	var res = dfs(s, words, []string{})
	var sentences []string
	for _, ss := range res {
		sentences = append(sentences, strings.Join(ss, " "))
	}
	return sentences
}

func main() {
	// 	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// 	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	// 	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	// fmt.Println(strings.Join(wordBreak("aaaaaaaa", []string{"aa", "aaa", "aaaa"}), "|"))
	// fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

	var s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	var wd = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	var res = wordBreak(s, wd)
	for _, s := range res {
		fmt.Println(s)
	}
}
