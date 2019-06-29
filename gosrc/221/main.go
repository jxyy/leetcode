package main

import "fmt"

func max(nums ...int) int {
	var m = 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func min(nums ...int) int {
	var m = 1 << 31
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

func maximalSquare(matrix [][]byte) int {
	var row = len(matrix)
	if row == 0 {
		return 0
	}
	var col = len(matrix[0])
	var dp = make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	var maxSize = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var up, left, upleft int
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 && j > 0 {
				upleft = dp[i-1][j-1]
			}

			if matrix[i][j] == '1' {
				dp[i][j] = min(up, left, upleft) + 1
				maxSize = max(dp[i][j], maxSize)
			}
		}
	}
	return maxSize * maxSize
}

func main() {
	var m = [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(m))
}
