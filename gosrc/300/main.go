package main

import "fmt"

func max(nums []int) int {
	var m = nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func lengthOfLIS(nums []int) int {
	var l = len(nums)
	if l == 0 {
		return 0
	}
	var dp = make([]int, l)
	for i := l - 1; i >= 0; i-- {
		var dpi = 1
		for j := i + 1; j < l; j++ {
			if nums[j] > nums[i] && dp[j]+1 > dpi {
				dpi = dp[j] + 1
			}
		}
		dp[i] = dpi
	}
	return max(dp)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{-2, -1}))
}
