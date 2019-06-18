package main

import "fmt"

func minThree(nums []int) int {
	var min = 0x7ffffffff
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	var lo, hi = 0, len(nums)
	for {
		var mi = (lo + hi) / 2
		if nums[mi] < nums[mi-1] {
			return nums[mi]
		} else if nums[mi] > nums[0] {
			lo = mi
		} else {
			hi = mi + 1
		}
	}
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{0, 1, 2}))
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{1}))
}
