package main

import (
	"fmt"
	"sort"
)

func findKthLargest0(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func partition(nums []int) int {
	var privot = nums[0]
	var lo, hi = 1, len(nums)
	for hi-lo > 0 {
		// fmt.Println(lo, hi, nums)
		var nlo = nums[lo]
		if nlo <= privot {
			nums[lo] = privot
			nums[lo-1] = nlo
			lo++
		} else {
			nums[lo] = nums[hi-1]
			nums[hi-1] = nlo
			hi--
		}
	}
	// fmt.Println(nums, privot, lo-1)
	return lo - 1
}

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	var lo, hi = 0, len(nums)
	var pk = partition(nums)
	for pk != k {
		// fmt.Println("kpk", k, pk)
		if pk < k {
			lo = pk + 1
			pk = lo + partition(nums[lo:hi])
		} else {
			hi = pk
			pk = lo + partition(nums[lo:hi])
		}
	}
	return nums[k]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
