package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Nums []string

func (t Nums) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Nums) Less(i, j int) bool {
	var si, sj = t[i], t[j]
	var li, lj = len(si), len(sj)
	var swaped = false
	if li > lj {
		si, sj = sj, si
		li, lj = lj, li
		swaped = true
	}
	if !strings.HasPrefix(sj, si) {
		if swaped {
			return !(si < sj)
		}
		return si < sj
	} else {
		if swaped {
			return !(sj < sj[li:]+si)
		}
		return sj < sj[li:]+si
	}
}

func (t Nums) Len() int {
	return len(t)
}

func largestNumber(nums []int) string {
	var strNums = make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}
	sort.Sort(sort.Reverse(Nums(strNums)))
	if strNums[0] == "0" {
		return "0"
	}
	return strings.Join(strNums, "")
}

func main() {
	// fmt.Println(largestNumber([]int{10, 2}))
	// fmt.Println(largestNumber([]int{3, 30, 5, 34, 9}))
	// fmt.Println(largestNumber([]int{1, 20}))
	// fmt.Println(largestNumber([]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}))
	fmt.Println(largestNumber([]int{3, 34}))
}
