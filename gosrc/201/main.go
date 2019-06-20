package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	var w = uint(0)
	for ; m != n; w++ {
		m = m >> 1
		n = n >> 1
	}
	return m << w
}

func main() {
	fmt.Println(rangeBitwiseAnd(1024, 2047))
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 1))
	fmt.Println(rangeBitwiseAnd(3, 3))
	fmt.Println(rangeBitwiseAnd(6, 7))
}
