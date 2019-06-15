package main

import "fmt"

func singleNumber(nums []int) []int {
	var a = 0
	for _, num := range nums {
		a ^= num
	}
	var mask = a & ^(a - 1)
	var s1, s2 = 0, 0
	for _, num := range nums {
		if mask&num != 0 {
			s1 ^= num
		} else {
			s2 ^= num
		}
	}
	return []int{s1, s2}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
