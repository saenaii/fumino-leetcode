package classic_problems

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
func isPalindrome(head *util.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for slow.Next != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverseLinkedList(slow)

	for head != nil && second != nil {
		if head.Val != second.Val {
			return false
		}
		head, second = head.Next, second.Next
	}

	return true
}

func reverseLinkedList(head *util.Node) *util.Node {
	newHead := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = newHead
		newHead = next
	}
	return newHead
}
