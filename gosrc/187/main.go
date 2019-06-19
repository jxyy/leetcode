package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	var count = make(map[string]bool)
	var l = len(s) + 1
	var res []string
	for i := 0; i < l-10; i++ {
		var ss = s[i : i+10]
		counted, meeted := count[ss]
		if !meeted {
			count[ss] = false
		} else if !counted {
			res = append(res, ss)
			count[ss] = true
		}
	}
	return res
}

func main() {
	// fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAA"))
}
