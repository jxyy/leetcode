package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	var l = len(s)
	for i := 0; i < l-i-1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func dfs(s string, curr []string, res [][]string) [][]string {
	var l = len(s)
	for i := 1; i < l+1; i++ {
		if isPalindrome(s[:i]) {
			var newone = make([]string, len(curr))
			copy(newone, curr)
			newone = append(newone, s[:i])
			if i == l {
				res = append(res, newone)
			} else {

				res = dfs(s[i:], newone, res)
			}
		}
	}
	return res
}

func partition(s string) [][]string {
	var res = dfs(s, []string{}, [][]string{})
	return res
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("aba"))
	fmt.Println(partition("aabbbaa"))
}
