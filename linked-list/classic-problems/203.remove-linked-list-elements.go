package classic_problems

import "fumino-leetcode/linked-list/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *util.Node, val int) *util.Node {
	if head == nil {
		return head
	}

	newHead := &util.Node{0, head}
	prev := newHead
	for prev.Next != nil {
		current := prev.Next
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = prev.Next
		}
	}
	return newHead.Next
}
