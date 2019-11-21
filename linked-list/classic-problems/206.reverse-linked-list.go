package classic_problems

import "fumino-leetcode/linked-list/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *util.Node) *util.Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = newHead
		newHead = next
	}
	return newHead
}
