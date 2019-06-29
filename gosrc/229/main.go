package main

import "fmt"

func majorityElement(nums []int) []int {
	var num1, num2, count1, count2 int
	for _, num := range nums {
		if count1 == 0 && num != num2 {
			num1 = num
		} else if count2 == 0 && num != num1 {
			num2 = num
		}
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		} else {
			count1--
			count2--
		}
		// fmt.Println(num1, count1, num2, count2)
	}
	// fmt.Println(num1, num2)

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}
	var res = []int{}
	var threshold = len(nums) / 3
	if count1 > threshold {
		res = append(res, num1)
	}
	if count2 > threshold {
		res = append(res, num2)
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
	fmt.Println(majorityElement([]int{1, 2, 2, 3, 2, 1, 1, 3}))
}
