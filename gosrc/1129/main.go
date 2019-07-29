package main

import "fmt"

type fromStatus struct {
	i        int
	fromRed  bool
	fromBlue bool
}

type visiteStatus struct {
	visitedFromRed  bool
	visitedFromBlue bool
}

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	var shortests = make([]int, n)
	for i := range shortests {
		shortests[i] = -1
	}
	shortests[0] = 0

	// prepare graph
	var rg, bg = make([][]bool, n), make([][]bool, n)
	for i := range shortests {
		rg[i] = make([]bool, n)
		bg[i] = make([]bool, n)
	}
	for _, edge := range red_edges {
		var src, dst = edge[0], edge[1]
		rg[src][dst] = true
	}
	for _, edge := range blue_edges {
		var src, dst = edge[0], edge[1]
		bg[src][dst] = true
	}

	var status = make([]visiteStatus, n)
	var pending, nextPending = []fromStatus{}, []fromStatus{fromStatus{0, false, false}}
	var step = 1
	for len(nextPending) > 0 {
		pending, nextPending = nextPending, pending
		for _, fs := range pending {
			if !fs.fromRed {
				for rdst, isConnected := range rg[fs.i] {
					if !isConnected {
						continue
					}
					if shortests[rdst] == -1 {
						shortests[rdst] = step
					}
					if !status[rdst].visitedFromRed {
						status[rdst].visitedFromRed = true
						nextPending = append(nextPending, fromStatus{rdst, true, false})
					}
				}
			}
			if !fs.fromBlue {
				for bdst, isConnected := range bg[fs.i] {
					if !isConnected {
						continue
					}
					if shortests[bdst] == -1 {
						shortests[bdst] = step
					}
					if !status[bdst].visitedFromBlue {
						status[bdst].visitedFromBlue = true
						nextPending = append(nextPending, fromStatus{bdst, false, true})
					}
				}
			}

		}
		step++
		pending = pending[:0]
	}

	return shortests
}

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{[]int{0, 1}, []int{1, 2}}, nil))
	fmt.Println(shortestAlternatingPaths(3, [][]int{[]int{0, 1}}, [][]int{[]int{2, 1}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{[]int{1, 0}}, [][]int{[]int{2, 1}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{[]int{1, 0}}, [][]int{[]int{1, 2}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{[]int{0, 1}, []int{0, 2}}, [][]int{[]int{1, 0}}))
	fmt.Println(shortestAlternatingPaths(5, [][]int{
		[]int{0, 1},
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
	}, [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 1},
	}))
}
