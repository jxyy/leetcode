package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func split(head *ListNode) (h1, h2 *ListNode) {
	var quick, slow = head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	h1 = head
	h2 = slow.Next
	slow.Next = nil
	return h1, h2
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var h1, h2 = split(head)
	h1 = sortList(h1)
	h2 = sortList(h2)
	var newHead, newTail *ListNode
	for h1 != nil || h2 != nil {
		var newNode *ListNode
		if h1 == nil {
			newNode = h2
			h2 = h2.Next
		} else if h2 == nil {
			newNode = h1
			h1 = h1.Next
		} else {
			if h1.Val < h2.Val {
				newNode = h1
				h1 = h1.Next
			} else {
				newNode = h2
				h2 = h2.Next
			}
		}

		if newHead == nil {
			newHead = newNode
			newTail = newNode
		} else {
			newTail.Next = newNode
			newTail = newNode
		}
	}
	newTail.Next = nil
	return newHead
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
	// var n5 = &ListNode{0, nil}
	var n4 = &ListNode{3, nil}
	var n3 = &ListNode{1, n4}
	var n2 = &ListNode{2, n3}
	var n1 = &ListNode{4, n2}
	print(n1)
	var n = sortList(n1)
	print(n)
}
