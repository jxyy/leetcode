package main

import "fmt"

type Status struct {
	max          int
	includeFirst bool
}

func dp(nums []int) int {
	var status1 = Status{nums[0], true}
	var status2 = Status{0, false}
	var l = len(nums)
	for i, num := range nums[1:] {
		var newStatus1, newStatus2 Status
		if i != l-2 || !status2.includeFirst {
			newStatus1 = Status{status2.max + num, status2.includeFirst}
		} else {
			newStatus1 = status2
		}

		if status1.max > status2.max || (status1.max == status2.max && !status1.includeFirst) {
			newStatus2.max = status1.max
			newStatus2.includeFirst = status1.includeFirst
		} else {
			newStatus2.max = status2.max
			newStatus2.includeFirst = status2.includeFirst
		}
		status1, status2 = newStatus1, newStatus2
	}
	if status1.max > status2.max {
		return status1.max
	}
	return status2.max
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var m = dp(nums)
	var m2 = dp(append(nums[1:], nums[0]))
	if m > m2 {
		return m
	}
	return m2
}

func main() {
	// fmt.Println(rob([]int{2, 3, 2}))
	// fmt.Println(rob([]int{2, 5, 2}))
	// fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3}))
}
