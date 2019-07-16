package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	var row = len(matrix)
	if row == 0 {
		return false
	}
	var col = len(matrix[0])
	if col == 0 {
		return false
	}

	var i, j = 0, col - 1
	for i < row && j >= 0 {
		var n = matrix[i][j]
		if n == target {
			return true
		} else if n > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	var m = [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(m, 5))
	fmt.Println(searchMatrix(m, 20))
}
