package classic_problems

import "fumino-leetcode/linked-list/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *util.Node) *util.Node {
	if head == nil || head.Next == nil {
		return head
	}

	even, odds := head, head.Next
	p, q := even, odds
	for p.Next != nil && q.Next != nil {
		p.Next = q.Next
		p = p.Next

		q.Next = p.Next
		q = q.Next
	}
	p.Next = odds
	return even
}
