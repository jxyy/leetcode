package main

import "fmt"

func min(n2, n3, n5 int) (i, n int) {
	i, n = 2, n2
	if n3 < n {
		i, n = 3, n3
	}
	if n5 < n {
		i, n = 5, n5
	}
	return i, n
}

func nthUglyNumber(n int) int {
	var nums = make([]int, 1, n)
	nums[0] = 1
	var i2, i3, i5 = 0, 0, 0
	for len(nums) < n {
		var i, minn = min(nums[i2]*2, nums[i3]*3, nums[i5]*5)
		switch i {
		case 2:
			i2++
		case 3:
			i3++
		case 5:
			i5++
		}
		if minn > nums[len(nums)-1] {
			nums = append(nums, minn)
		}
	}
	return nums[n-1]
}

func main() {
	for i := 0; i < 30; i++ {
		fmt.Println(i+1, nthUglyNumber(i+1))
	}
}
