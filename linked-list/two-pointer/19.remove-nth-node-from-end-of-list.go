package two_pointer

import (
	"fumino-leetcode/linked-list/util"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *util.Node, n int) *util.Node {
	if head == nil || n == 0 {
		return head
	}

	prev, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast != nil && fast.Next != nil {
		prev = prev.Next
		fast = fast.Next
	}

	prev.Next = prev.Next.Next

	return head
}
