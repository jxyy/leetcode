package main

import "fmt"

func longestWPI(hours []int) int {
	var lwpi = 0
	var sum = 0
	var first = make(map[int]int)
	for i, h := range hours {
		var v = -1
		if h > 8 {
			v = 1
		}
		sum += v
		if sum > 0 {
			lwpi = i + 1
		} else {
			if fi, ok := first[sum-1]; ok && i-fi > lwpi {
				lwpi = i - fi
			}
			if _, ok := first[sum]; !ok {
				first[sum] = i
			}
		}
	}
	return lwpi
}

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	fmt.Println(longestWPI([]int{6, 6, 9}))
	// fmt.Println([]int{})
}
