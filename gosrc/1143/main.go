package main

import "fmt"

type status struct {
	i     int
	j     int
	count int
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var l1, l2 = len(text1), len(text2)
	var dp = make([][]status, l1)
	for i := range dp {
		dp[i] = make([]status, l2)
		for j := range dp[i] {
			dp[i][j] = status{-1, -1, 0}
		}
	}

	for i := 0; i < l1; i++ {
		var c1, c2 = text1[i], byte(' ')
		for j := 0; j < l2; j++ {
			c2 = text2[j]
			var up, left = status{-1, -1, 0}, status{-1, -1, 0}
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}

			if c1 == c2 && i > up.i && j > up.j {
				up = status{i, j, up.count + 1}
			}
			if c1 == c2 && i > left.i && j > left.j {
				left = status{i, j, left.count + 1}
			}

			if up.count > left.count {
				dp[i][j] = up
			} else {
				dp[i][j] = left
			}

		}
	}
	// fmt.Println(dp)
	return dp[l1-1][l2-1].count
}

func main() {
	fmt.Println(longestCommonSubsequence("abced", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
}
