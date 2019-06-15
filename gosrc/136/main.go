package main

import "fmt"

func singleNumber(nums []int) int {
	var a = 0
	for _, num := range nums {
		a ^= num
	}
	return a
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
