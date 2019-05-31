package main

import (
	"fmt"
)

func isEqual(c1, c2 byte) bool {
	return c1 == '?' || c2 == '?' || c1 == c2
}

func find(s, sub string) int {
	var next = make([]int, len(sub))
	var l = 0
	var i = 1
	next[0] = 0
	for i < len(sub) {
		if sub[l] == sub[i] {
			l++
			next[i] = l
			i++
		} else {
			for ; l > 0 && sub[l] != sub[i]; l = next[l-1] {
			}
			next[i] = l
			i++
		}
	}
	// fmt.Println("prepare finish", next)

	for i, j := 0, 0; i < len(s); {
		if s[i] == sub[j] {
			i++
			j++
			if j == len(sub) {
				return i
			}
		} else {
			for ; j > 0 && s[i] != sub[j]; j = next[j-1] {
			}
			if j == 0 && s[i] != sub[j] {
				i++
			}
		}
	}
	return -1
}

func isMatch(s string, p string) bool {
	// fmt.Println("\""+s+"\"", "\""+p+"\"")

	var ls, lp = len(s), len(p)
	var i, j = 0, 0
	var iStart, jStart = -1, -1
	for i < ls {
		if j < lp && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < lp && p[j] == '*' {
			j++
			jStart = j
			iStart = i
		} else if jStart != -1 {
			j = jStart
			iStart++
			i = iStart
		} else {
			return false
		}
	}

	for ; j < lp && p[j] == '*'; j++ {
	}
	return i == ls && j == lp
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("acdcb", "*a*b"))
	fmt.Println(isMatch("aa", "a*c?b"))

	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("aaaa", "***a"))
	fmt.Println(isMatch("", "*"))
	fmt.Println(isMatch("aac", "*ac"))
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}
