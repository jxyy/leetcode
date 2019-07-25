package main

type NumArray2 struct {
	nums      []int
	prefixSum []int
}

func Constructor2(nums []int) NumArray2 {
	var psum = make([]int, len(nums))
	var sum = 0
	for i, num := range nums {
		sum += num
		psum[i] = sum
	}
	return NumArray2{nums, psum}
}

func (this *NumArray2) Update(i int, val int) {
	var diff = val - this.nums[i]
	for j := i; j < len(this.prefixSum); j++ {
		this.prefixSum[j] += diff
	}
	this.nums[i] = val
	// fmt.Println(this.nums, this.prefixSum)
}

func (this *NumArray2) SumRange(i int, j int) int {
	var n1, n2 int
	if i > 0 {
		n1 = this.prefixSum[i-1]
	}
	n2 = this.prefixSum[j]
	return n2 - n1
}
