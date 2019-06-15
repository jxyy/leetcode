package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var l []*ListNode
	var node = head
	for node != nil {
		l = append(l, node)
		node = node.Next
	}

	var n = len(l)
	node = nil
	for i, j := 0, n-1; i < j; {
		var ni, nj = l[i], l[j]
		if node != nil {
			node.Next = ni
		}
		ni.Next = nj
		node = nj
		i++
		j--
	}
	if n%2 == 1 {
		node.Next = l[n/2]
		node = node.Next
	}
	node.Next = nil
}

func print(n *ListNode) {
	var node = n
	for node != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// var n5 = &ListNode{5, nil}
	var n4 = &ListNode{4, nil}
	var n3 = &ListNode{3, n4}
	var n2 = &ListNode{2, n3}
	var n1 = &ListNode{1, n2}
	print(n1)
	reorderList(n1)
	print(n1)
}
