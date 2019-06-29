package main

import "fmt"

func intersection(min1, max1, min2, max2 int) int {
	var left = min1
	if min2 > left {
		left = min2
	}
	var right = max1
	if max2 < right {
		right = max2
	}
	var d = right - left
	if d < 0 {
		return 0
	}
	return d
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	var s1 = (C - A) * (D - B)
	var s2 = (G - E) * (H - F)
	var coveredWidth = intersection(A, C, E, G)
	var coveredHeight = intersection(B, D, F, H)
	return s1 + s2 - coveredWidth*coveredHeight
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
