package main

import "fmt"

type NumArray struct {
	n    int
	bt   []int
	nums []int
}

func Constructor(nums []int) NumArray {
	var n = len(nums)
	var bt = make([]int, n+1)
	var na = NumArray{n, bt, make([]int, n)}
	for i, val := range nums {
		na.Update(i, val)
	}
	na.nums = nums
	return na
}

func (this *NumArray) Update(i int, val int) {
	var df = val - this.nums[i]
	this.nums[i] = val
	i++
	for i <= this.n {
		this.bt[i] += df
		i += i & (-i)
	}
}

func (this *NumArray) sum(i int) int {
	var s = 0
	i++
	for i > 0 {
		s += this.bt[i]
		i -= i & (-i)
	}
	return s
}

func (this *NumArray) SumRange(i int, j int) int {
	var n1, n2 int
	if i > 0 {
		n1 = this.sum(i - 1)
	}
	n2 = this.sum(j)
	return n2 - n1
}

func main() {
	var na = Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))
}
