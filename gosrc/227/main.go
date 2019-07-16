package main

import "fmt"

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func calculate(s string) int {
	var nums []int
	var ops []byte
	var num = 0
	for i, c := range s {
		switch c {
		case '+', '-':
			ops = append(ops, byte(c))
		case '*', '/':
			ops = append(ops, byte(c))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num *= 10
			num += int(c - '0')
			if (i < len(s)-1 && !isDigit(s[i+1])) || i == len(s)-1 {
				nums = append(nums, num)
				num = 0
				var lo, ln = len(ops), len(nums)
				if lo > 0 {
					var lastOp = ops[lo-1]
					switch lastOp {
					case '*':
						num1, num2 := nums[ln-2], nums[ln-1]
						nums[ln-2] = num1 * num2
						ops = ops[:lo-1]
						nums = nums[:ln-1]
					case '/':
						num1, num2 := nums[ln-2], nums[ln-1]
						nums[ln-2] = num1 / num2
						ops = ops[:lo-1]
						nums = nums[:ln-1]
					}
				}
			}
		}
	}

	//eval rest part
	var io, in = 0, 0
	var lo = len(ops)
	for io < lo {
		var op = ops[io]
		var num1, num2 = nums[in], nums[in+1]
		var newNum int
		switch op {
		case '+':
			newNum = num1 + num2
		case '-':
			newNum = num1 - num2
		}
		nums[in+1] = newNum
		io++
		in++
	}
	return nums[in]
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3 /2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
	fmt.Println(calculate("1-1+1"))
}
