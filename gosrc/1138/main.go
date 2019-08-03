package main

import "fmt"

type position struct {
	i int
	j int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func alphabetBoardPath(target string) string {
	var charPos = make(map[byte]position)
	for i := 0; i < 26; i++ {
		var c = byte('a' + i)
		charPos[c] = position{i / 5, i % 5}
	}

	var ops []byte
	var i, j = 0, 0
	for _, c := range target {
		var cpos = charPos[byte(c)]
		var di, dj = cpos.i - i, cpos.j - j
		// fmt.Printf("%c %d %d %d %d\n", c, di, dj, i, j)
		if di > 0 {
			// move down
			for k := 0; k < abs(dj); k++ {
				if dj > 0 {
					ops = append(ops, 'R')
				} else {
					ops = append(ops, 'L')
				}
			}

			for k := 0; k < abs(di); k++ {
				if di > 0 {
					ops = append(ops, 'D')
				} else {
					ops = append(ops, 'U')
				}
			}
		} else {
			for k := 0; k < abs(di); k++ {
				if di > 0 {
					ops = append(ops, 'D')
				} else {
					ops = append(ops, 'U')
				}
			}

			for k := 0; k < abs(dj); k++ {
				if dj > 0 {
					ops = append(ops, 'R')
				} else {
					ops = append(ops, 'L')
				}
			}
		}

		ops = append(ops, '!')
		i, j = cpos.i, cpos.j
	}
	return string(ops)
}

func main() {
	fmt.Println(alphabetBoardPath("leet"))
	fmt.Println(alphabetBoardPath("code"))
	fmt.Println(alphabetBoardPath("zdz"))
}
