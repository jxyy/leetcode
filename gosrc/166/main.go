package main

import (
	"fmt"
	"strconv"
)

func digitsToString(digits []int) string {
	var bts = make([]byte, len(digits))
	for i, d := range digits {
		bts[i] = byte(d) + 48
	}
	return string(bts)
}

func fractionToDecimal(numerator int, denominator int) string {
	if denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}

	var isNegtive = false
	if numerator < 0 {
		isNegtive = true
		numerator = -numerator
	}

	var intPart = numerator / denominator
	var residual = numerator % denominator
	if residual == 0 {
		if isNegtive {
			return "-" + strconv.Itoa(intPart)
		}
		return strconv.Itoa(intPart)
	}

	var decimalPart []int
	var mem = make(map[int]int)
	var loopPos = -1
	mem[residual] = 0
	for i := 1; residual > 0; i++ {
		residual *= 10
		decimalPart = append(decimalPart, residual/denominator)
		residual = residual % denominator
		if pos, ok := mem[residual]; ok {
			loopPos = pos
			break
		}
		mem[residual] = i
	}

	var decimalStr string
	if loopPos >= 0 {
		decimalStr = digitsToString(decimalPart[:loopPos]) + "(" + digitsToString(decimalPart[loopPos:]) + ")"
	} else {
		decimalStr = digitsToString(decimalPart)
	}
	if isNegtive {
		return "-" + strconv.Itoa(intPart) + "." + decimalStr
	}
	return strconv.Itoa(intPart) + "." + decimalStr
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(-50, 8))
	fmt.Println(fractionToDecimal(1, 214748364))
}
