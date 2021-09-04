// Definition for singly-linked list.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	// var dummy *ListNode = new(ListNode)
	// dummy : = new(ListNode)
	head := dummy
	for !(l1 == nil && l2 == nil) {
		if l1 != nil && (l2 == nil || l1.Val < l2.Val) {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}
	return head.Next
}
