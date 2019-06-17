package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var maxProduct = 0
	var product = 1
	var negPrefix = 1
	for _, num := range nums {
		product *= num
		if product == 0 {
			product = 1
			negPrefix = 1
		} else if product < 0 {
			if negPrefix > 0 {
				negPrefix = product
			} else {
				maxProduct = max(maxProduct, product/negPrefix)
			}
		} else {
			maxProduct = max(maxProduct, product)
		}
	}
	return maxProduct
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-4, -3}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	fmt.Println(maxProduct([]int{3, -1, 4}))
}
