package main

import "fmt"

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	var skillMap = make(map[string]int)
	for i, skill := range req_skills {
		skillMap[skill] = i
	}

	var dp = make(map[int][]int)
	dp[0] = []int{}
	for i, pSkills := range people {
		var pstatus = 0
		for _, skill := range pSkills {
			pstatus |= 1 << uint(skillMap[skill])
		}

		for prevStatus, prevPeoples := range dp {
			var nextStatus = pstatus | prevStatus
			var existPeoples, ok = dp[nextStatus]
			if !ok || len(existPeoples) > len(prevPeoples)+1 {
				var newPeoples = make([]int, len(prevPeoples))
				copy(newPeoples, prevPeoples)
				dp[nextStatus] = append(newPeoples, i)
			}
		}
	}
	return dp[(1<<uint(len(req_skills)))-1]
}

func main() {
	fmt.Println(smallestSufficientTeam(
		[]string{"java", "nodejs", "reactjs"},
		[][]string{
			[]string{"java"},
			[]string{"nodejs"},
			[]string{"nodejs", "reactjs"},
		},
	))

	fmt.Println(smallestSufficientTeam(
		[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{
			[]string{"algorithms", "math", "java"},
			[]string{"algorithms", "math", "reactjs"},
			[]string{"java", "csharp", "aws"},
			[]string{"reactjs", "csharp"},
			[]string{"csharp", "math"},
			[]string{"aws", "java"},
		},
	))
}
