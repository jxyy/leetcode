package main

import "fmt"

func numIslands(grid [][]byte) int {
	var rowLen = len(grid)
	if rowLen == 0 {
		return 0
	}
	var colLen = len(grid[0])
	var visited = make([][]bool, len(grid))
	var pendingi, pendingj []int
	var landCellCount = 0
	for i, row := range grid {
		visited[i] = make([]bool, len(row))
		for _, cell := range row {
			if cell == '1' {
				landCellCount++
			}
		}
	}

	var visitedCount = 0
	var landCount = 0
	for visitedCount < landCellCount {
		// fmt.Println("start", visitedCount, landCellCount)
		for i := 0; i < rowLen && len(pendingi) == 0; i++ {
			for j := range grid[i] {
				if grid[i][j] == '1' && !visited[i][j] {
					pendingi, pendingj = append(pendingi, i), append(pendingj, j)
					visited[i][j] = true
					break
				}
			}
		}
		// fmt.Println("mid", len(pendingi), len(pendingj))
		for len(pendingi) > 0 {
			var l = len(pendingi)
			var i, j = pendingi[l-1], pendingj[l-1]
			pendingi, pendingj = pendingi[:l-1], pendingj[:l-1]
			// visited[i][j] = true
			// fmt.Println("visit", i, j, l)
			visitedCount++
			if i > 0 && grid[i-1][j] == '1' && !visited[i-1][j] {
				pendingi, pendingj = append(pendingi, i-1), append(pendingj, j)
				visited[i-1][j] = true
			}
			if i+1 < rowLen && grid[i+1][j] == '1' && !visited[i+1][j] {
				pendingi, pendingj = append(pendingi, i+1), append(pendingj, j)
				visited[i+1][j] = true
			}
			if j > 0 && grid[i][j-1] == '1' && !visited[i][j-1] {
				pendingi, pendingj = append(pendingi, i), append(pendingj, j-1)
				visited[i][j-1] = true
			}
			if j+1 < colLen && grid[i][j+1] == '1' && !visited[i][j+1] {
				pendingi, pendingj = append(pendingi, i), append(pendingj, j+1)
				visited[i][j+1] = true
			}
		}
		landCount++
		// fmt.Println("end", visitedCount, landCellCount)
	}
	return landCount
}

func main() {
	var grid = [][]byte{
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '1', '1'},
		[]byte{'0', '0', '1', '1', '0'},
		[]byte{'0', '0', '0', '0', '1'},
	}

	fmt.Println(numIslands(grid))
}
