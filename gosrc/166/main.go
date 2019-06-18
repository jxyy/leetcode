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
	var intPart = numerator / denominator
	var residual = numerator % denominator
	if residual == 0 {
		return strconv.Itoa(intPart)
	}

	var decimalPart []int
	var mem = make([]int, denominator)
	for i := 0; i < denominator; i++ {
		mem[i] = -1
	}
	var loopPos = -1
	mem[residual] = 0
	for i := 0; residual > 0; i++ {
		residual *= 10
		decimalPart = append(decimalPart, residual/denominator)
		residual = residual % denominator
		loopPos = mem[residual]
		if loopPos >= 0 {
			break
		}
	}

	var decimalStr string
	if loopPos >= 0 {
		decimalStr = digitsToString(decimalPart[:loopPos]) + "(" + digitsToString(decimalPart[loopPos:]) + ")"
	} else {
		decimalStr = digitsToString(decimalPart)
	}
	return strconv.Itoa(intPart) + "." + decimalStr
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(digitsToString([]int{1, 2, 3}))
	fmt.Println(byte('0'))
}
