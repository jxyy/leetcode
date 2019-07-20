package main

import (
	"fmt"
)

func hIndex(citations []int) int {
	// sorted acending
	var l = len(citations)
	var lo, hi = 0, l
	for lo < hi {
		var mi = (lo + hi) / 2
		if citations[mi] >= l-mi {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return l - lo
}

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
}
