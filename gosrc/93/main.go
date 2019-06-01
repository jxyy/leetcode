package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	var ls = len(s)
	switch ls {
	case 1:
		return true
	case 2:
		return s[0] != '0'
	case 3:
		return s[0] != '0' && s <= "255"
	default:
		return false
	}
}

func search(s string, depth int, curr []string, result []string) []string {
	// fmt.Println(s, depth, curr, result)
	if depth == 0 || s == "" {
		if depth == 0 && s == "" {
			// fmt.Println("bingo")
			result = append(result, strings.Join(curr, "."))
		}
		return result
	}

	for i := 1; i <= len(s) && i <= 3; i++ {
		if isValid(s[:i]) {
			curr[4-depth] = s[:i]
			result = search(s[i:], depth-1, curr, result)
			curr[4-depth] = ""
		}
	}
	return result
}

func restoreIpAddresses(s string) []string {
	var result []string
	var curr = make([]string, 4, 4)
	return search(s, 4, curr, result)
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("010010"))
}
