package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var revSize = n - m + 1
	var stack = make([]*ListNode, revSize)
	var si = 0
	var i = 1
	var curr = head
	var nodep, nodeq *ListNode
	for curr != nil && i <= n {
		if i >= m && i <= n {
			stack[si] = curr
			si++
		}
		if i == n {
			nodeq = curr.Next
		}
		i++
		if i == m {
			nodep = curr
		}
		curr = curr.Next
	}

	if nodep == nil {
		head = stack[revSize-1]
	} else {
		nodep.Next = stack[revSize-1]
	}
	var prev *ListNode
	curr = nil
	for i = revSize - 1; i >= 0; i-- {
		curr = stack[i]
		if prev != nil {
			prev.Next = curr
		}
		prev = curr
	}
	if curr != nil {
		curr.Next = nodeq
	}
	return head
}

func (head *ListNode) print() {
	var curr = head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Print("nil\n")
}

func newList(nums []int) (head *ListNode) {
	var prev *ListNode
	for _, num := range nums {
		curr := ListNode{num, nil}
		if prev != nil {
			prev.Next = &curr
		}
		if head == nil {
			head = &curr
		}
		prev = &curr
	}
	return head
}

func main() {
	var l = newList([]int{1, 2, 3, 4, 5})
	l.print()
	l = reverseBetween(l, 2, 4)
	l.print()
	l = reverseBetween(l, 1, 3)
	l.print()
	l = reverseBetween(l, 4, 5)
	l.print()
}
