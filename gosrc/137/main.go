package main

import "fmt"

func singleNumber(nums []int) int {
	var b1, b2 = 0, 0
	for _, num := range nums {
		b1 = (num & ^b1 & ^b2) | (^num & b1)
		b2 = (^b2 & num & ^b1) | (^num & b2)
	}
	return b1
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
