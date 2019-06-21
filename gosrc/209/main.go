package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	var l = len(nums)
	if l == 0 {
		return 0
	}

	var i, j = 0, 1
	var minl = l + 1
	var sum = nums[0]
	for i <= j && j < l {
		for j < l && sum < s {
			sum += nums[j]
			j++
		}
		if sum >= s {
			minl = min(minl, j-i)
		}
		for i <= j && sum-nums[i] >= s {
			sum -= nums[i]
			i++
		}
		if sum >= s {
			minl = min(minl, j-i)
		}

		sum -= nums[i]
		i++
	}
	if minl < l+1 {
		return minl
	}
	return 0
}

func main() {
	// fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
}
