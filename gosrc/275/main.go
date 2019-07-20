package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		var dpi = i
		for j := 1; j*j <= i; j++ {
			dpi = min(dpi, dp[i-j*j]+1)
		}
		dp[i] = dpi
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}
