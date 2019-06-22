package main

import "fmt"

// func dfs(k, n, i, currSum int, curr []int, res [][]int) [][]int {
// 	if i == k-1

// 	return res
// }

func combinationSum3(k int, n int) [][]int {
	var curr = make([]int, k)
	var currSum = 0
	var res [][]int
	var dfs func(int)
	dfs = func(i int) {
		var from = 0
		if i > 0 {
			from = curr[i-1]
		}

		fmt.Println(i, curr, from)
		if i == k-1 {
			var rest = n - currSum
			if rest > from && rest < 10 {
				curr[i] = rest
				var newCurr = make([]int, k)
				copy(newCurr, curr)
				res = append(res, newCurr)
			}
			return
		}
		for j := from + 1; j < 10; j++ {
			if currSum+j < n {
				curr[i] = j
				currSum += j
				dfs(i + 1)
				currSum -= j
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(combinationSum3(2, 18))
}
