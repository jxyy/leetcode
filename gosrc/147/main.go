package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var sortedHead, sortedTail = head, head
	var unsoredHead = head.Next
	for unsoredHead != nil {
		var nextUnsorted = unsoredHead.Next

		var nodeToInsert = unsoredHead
		if nodeToInsert.Val < sortedHead.Val {
			nodeToInsert.Next = sortedHead
			sortedHead = nodeToInsert
			sortedTail.Next = nextUnsorted
		} else if nodeToInsert.Val >= sortedTail.Val {
			// do nothing
			sortedTail = nodeToInsert
		} else {
			var node = sortedHead
			for ; node.Next.Val <= nodeToInsert.Val; node = node.Next {
			}
			var tmp = node.Next
			node.Next = nodeToInsert
			nodeToInsert.Next = tmp
			sortedTail.Next = nextUnsorted
		}

		unsoredHead = nextUnsorted

		// var n = sortedHead
		// for n != sortedTail.Next {
		// 	fmt.Print(n.Val, " -> ")
		// 	n = n.Next
		// }
		// fmt.Println()
		// bufio.NewReader(os.Stdin).ReadByte()
	}
	sortedTail.Next = nil
	return sortedHead
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
	var n5 = &ListNode{4, nil}
	var n4 = &ListNode{2, n5}
	var n3 = &ListNode{5, n4}
	var n2 = &ListNode{1, n3}
	var n1 = &ListNode{3, n2}
	print(n1)

	var n = insertionSortList(n1)
	print(n)
}
