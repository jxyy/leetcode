package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	var hindex = 0
	for _, c := range citations {
		if c > hindex {
			hindex++
		}
	}
	return hindex
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
