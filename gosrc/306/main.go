package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	var isAddtiive = false
	var l = len(num)
	var l1, l2 = 0, 0
	var dfs func(int)
	var nums = make([]int, 0, l)
	dfs = func(i int) {
		if isAddtiive {
			return
		}
		if i == l && len(nums) > 2 {
			isAddtiive = true
			return
		}

		var ll = len(nums)
		if ll == 0 {
			nums = append(nums, 0)
			if l > 0 && num[0] == '0' {
				l1 = 1
				dfs(1)
			} else {
				for l1 = 1; l1 <= l; l1++ {
					var n, _ = strconv.Atoi(num[:l1])
					nums[0] = n
					dfs(l1)
				}
			}
		} else if ll == 1 {
			nums = append(nums, 0)
			if l > l1 && num[l1] == '0' {
				l2 = 1
				dfs(l1 + 1)
			} else {
				for l2 = 1; l1+l2 <= l && num[l1] != '0'; l2++ {
					var n, _ = strconv.Atoi(num[l1 : l1+l2])
					nums[1] = n
					dfs(l1 + l2)
				}
			}
			nums = nums[:1]
		} else {
			var prev1, prev2 = nums[ll-2], nums[ll-1]
			var addon = prev1 + prev2
			var next = strconv.Itoa(addon)
			if strings.HasPrefix(num[i:], next) {
				nums = append(nums, addon)
				dfs(i + len(next))
				nums = nums[:len(nums)-1]
			}
		}
	}
	dfs(0)
	return isAddtiive
}

func main() {
	// fmt.Println(isAdditiveNumber("112358"))
	// fmt.Println(isAdditiveNumber("199100199"))
	// fmt.Println(isAdditiveNumber("1023"))
	// fmt.Println(isAdditiveNumber("101"))
	// fmt.Println(isAdditiveNumber("0"))
	fmt.Println(isAdditiveNumber("211738"))
}
