package main

import "fmt"

type pos struct {
	i int
	j int
}

func gameOfLife(board [][]int) {
	var row = len(board)
	if row == 0 {
		return
	}
	var col = len(board[0])

	var dirs = [][2]int{
		[2]int{1, 0},
		[2]int{1, 1},
		[2]int{0, 1},
		[2]int{1, -1},
		[2]int{0, -1},
		[2]int{-1, -1},
		[2]int{-1, 0},
		[2]int{-1, 1},
	}
	var update []pos
	for i, r := range board {
		for j, v := range r {
			var c = 0
			for _, dir := range dirs {
				var ii, jj = i + dir[0], j + dir[1]
				if ii >= 0 && ii < row && jj >= 0 && jj < col && board[ii][jj] == 1 {
					c++
				}
			}
			if v == 1 {
				if c < 2 || c > 3 {
					update = append(update, pos{i, j})
				}
			} else {
				if c == 3 {
					update = append(update, pos{i, j})
				}
			}

		}
	}
	for _, p := range update {
		var i, j = p.i, p.j
		board[i][j] = 1 - board[i][j]
	}
}

func main() {
	var board = [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}
