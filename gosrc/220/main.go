package main

import "fmt"

type Section struct {
	left  int
	right int
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 0 {
		return false
	}
	var sections = make([]*Section, k)
	for i, num := range nums {
		var ni = i % k
		for _, sec := range sections {
			if sec != nil && num >= sec.left && num <= sec.right {
				return true
			}
		}
		sections[ni] = &Section{num - t, num + t}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
