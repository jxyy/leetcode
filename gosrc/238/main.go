package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var l = len(nums)
	var out = make([]int, l)
	for i := range out {
		out[i] = 1
	}

	var s = 1
	for i, num := range nums {
		s *= num
		if i < l-1 {
			out[i+1] = s
		}
	}
	s = 1
	for i := range nums {
		var j = l - 1 - i
		var num = nums[j]
		s *= num
		if j > 0 {
			out[j-1] *= s
		}
	}
	return out
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{0, 0}))
}
