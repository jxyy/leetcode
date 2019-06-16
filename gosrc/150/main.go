package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var numStack []int
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			var ls = len(numStack)
			var num1 = numStack[ls-2]
			var num2 = numStack[ls-1]
			var newNum int
			switch token {
			case "+":
				newNum = num1 + num2
			case "-":
				newNum = num1 - num2
			case "*":
				newNum = num1 * num2
			case "/":
				newNum = num1 / num2
			}
			numStack[ls-2] = newNum
			numStack = numStack[:ls-1]
		default:
			var num, _ = strconv.Atoi(token)
			numStack = append(numStack, num)
		}
	}
	return numStack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
