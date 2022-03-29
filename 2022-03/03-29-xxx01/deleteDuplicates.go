package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	pre := head
	now := pre.Next
	for now != nil {
		if now.Val == pre.Val {
			for now = now.Next; now != nil; now = now.Next {
				if now.Val != pre.Val {
					break
				}
			}
			if pre == head {
				head = now
				first = now
				pre = now
			} else {
				first.Next = now
				pre = now
			}
			if pre != nil {
				now = pre.Next
			} else {
				break
			}
		} else {
			if pre != head {
				first = pre
			}
			pre = now
			now = now.Next
		}
	}
	return head
}
