package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var from = 0
	var prev = nums[0] - 1
	var intervals []string
	for i, num := range nums {
		if num != prev+1 {
			if i-1 == from {
				intervals = append(intervals, strconv.Itoa(nums[from]))
			} else {
				intervals = append(intervals, fmt.Sprintf("%d->%d", nums[from], nums[i-1]))
			}
			from = i
		}
		prev = num
	}
	if len(nums)-1 == from {
		intervals = append(intervals, strconv.Itoa(nums[from]))
	} else {
		intervals = append(intervals, fmt.Sprintf("%d->%d", nums[from], nums[len(nums)-1]))
	}
	return intervals
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
