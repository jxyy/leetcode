package main

import "fmt"

func distance(w1, w2 string) int {
	var c = 0
	for i := range w1 {
		// fmt.Println(w1[i], w2[i])
		if w1[i] != w2[i] {
			c++
		}
	}
	return c
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var graph = make([][]int, len(wordList))
	var endi = -1
	for i, w1 := range wordList {
		var l = []int{}
		for j, w2 := range wordList {
			if distance(w1, w2) == 1 {
				l = append(l, j)
			}
		}
		graph[i] = l
		if w1 == endWord {
			endi = i
		}
	}
	if endi == -1 {
		return 0
	}

	var depth = 1
	var isDone = false
	var pending []int
	var visited = make([]bool, len(wordList))
	for i, w := range wordList {
		if distance(beginWord, w) == 1 {
			pending = append(pending, i)
		}
	}

	for len(pending) > 0 {
		// for _, i := range pending {
		// 	fmt.Print(wordList[i], ", ")
		// }
		// fmt.Println("")

		depth++
		var nextPending = []int{}
		for _, i := range pending {
			if i == endi {
				isDone = true
				goto end
			}
			for _, j := range graph[i] {
				if !visited[j] {
					visited[j] = true
					nextPending = append(nextPending, j)
				}
			}
		}
		pending = nextPending
	}

end:
	if isDone {
		return depth
	} else {
		return 0
	}
}

func main() {
	// var bw, ew = "hit", "cog"
	// var wl = []string{"hot", "dot", "dog", "lot", "log", "cog"}

	var bw, ew = "hit", "cog"
	var wl = []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println(ladderLength(bw, ew, wl))

}
