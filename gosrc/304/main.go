package main

import "fmt"

type NumMatrix struct {
	cache [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var this = new(NumMatrix)
	this.cache = make([][]int, len(matrix))
	for i, line := range matrix {
		var cacheLine = make([]int, len(line))
		var lineSum = 0
		for j, num := range line {
			lineSum += num
			cacheLine[j] = lineSum
			if i > 0 {
				cacheLine[j] += this.cache[i-1][j]
			}
		}
		this.cache[i] = cacheLine
	}

	return *this
}

func (this *NumMatrix) debug() {
	fmt.Println(this.cache)
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var c1, c2, c3, c4 = 0, 0, 0, 0
	if row1 > 0 && col1 > 0 {
		c1 = this.cache[row1-1][col1-1]
	}
	if row1 > 0 {
		c2 = this.cache[row1-1][col2]
	}
	if col1 > 0 {
		c3 = this.cache[row2][col1-1]
	}
	c4 = this.cache[row2][col2]
	return c4 - c2 - c3 + c1
}

func main() {
	var nm = Constructor([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(nm.SumRegion(1, 1, 2, 2)) // 28
	fmt.Println(nm.SumRegion(0, 0, 1, 1)) // 12
	fmt.Println(nm.SumRegion(0, 1, 1, 2))
}
