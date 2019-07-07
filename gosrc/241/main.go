package main

import (
	"fmt"
)

func parse(input string) (nums []int, ops []byte) {
	var n = 0
	for _, c := range input {
		switch c {
		case '+', '-', '*':
			nums = append(nums, n)
			ops = append(ops, byte(c))
			n = 0
		default:
			n *= 10
			n += int(byte(c) - '0')
		}
	}
	nums = append(nums, n)
	return
}

func diffWaysToCompute(input string) []int {
	var nums, ops = parse(input)
	var ln = len(nums)
	var mem = make([][][]int, ln)
	for i := range mem {
		mem[i] = make([][]int, ln-i)
	}

	var eval func(int, int) []int
	eval = func(start int, l int) []int {
		// fmt.Println("in:", start, l)
		// bufio.NewReader(os.Stdin).ReadByte()
		if r := mem[start][l-1]; r != nil {
			return r
		}
		if l == 1 {
			mem[start][0] = []int{nums[start]}
			return mem[start][0]
		}
		var res []int
		for i := range ops[start : start+l-1] {
			var lstart, llen = start, i + 1
			var rstart, rlen = start + i + 1, l - i - 1
			var lres = eval(lstart, llen)
			var rres = eval(rstart, rlen)
			for _, n1 := range lres {
				for _, n2 := range rres {
					switch ops[start+i] {
					case '+':
						// fmt.Println("+", n1, n2)
						res = append(res, n1+n2)
					case '-':
						// fmt.Println("-", n1, n2)
						res = append(res, n1-n2)
					case '*':
						// fmt.Println("*", n1, n2)
						res = append(res, n1*n2)
					}
				}
			}
		}
		mem[start][l-1] = res
		return res
	}
	return eval(0, ln)
}

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
