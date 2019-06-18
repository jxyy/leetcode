package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func findPeakElement(nums []int) int {
	var lo, hi = 0, len(nums)
	// if hi == 1 {
	// 	return 0
	// }
	for {
		var mi = (hi + lo) / 2
		var n1 = MinInt
		var n2 = nums[mi]
		var n3 = MinInt
		if mi-1 >= 0 {
			n1 = nums[mi-1]
		}
		if mi+1 < len(nums) {
			n3 = nums[mi+1]
		}
		// fmt.Println(lo, hi, mi, n1, n2, n3)
		// fmt.Println(lo, hi, mi)
		if n1 < n2 && n2 > n3 {
			return mi
		} else if n1 > n2 {
			hi = mi
		} else {
			lo = mi
		}
	}
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1}))
	fmt.Println(findPeakElement([]int{1, 2}))
	fmt.Println(findPeakElement([]int{2, 1}))
	// fmt.Printf("%x ", MinInt)
}
