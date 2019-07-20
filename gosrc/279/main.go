package main

import (
	"fmt"
)

func hIndex(citations []int) int {
	// sorted acending
	var hindex = 0
	var l = len(citations)
	for i := 0; i < l; i++ {
		var c = citations[l-1-i]
		if c > hindex {
			hindex++
		}
	}
	return hindex
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
