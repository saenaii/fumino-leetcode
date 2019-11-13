package two_pointer

import "fumino-leetcode/linked-list/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *util.Node) *util.Node {
	if head == nil || head.Next == nil {
		return nil
	}

	if node := getCycleNode(head); node != nil {
		slow, fast := head, node.Next
		for slow != nil && fast != nil {
			if slow == fast {
				return slow
			}
			slow = slow.Next
			fast = fast.Next
		}
	}
	return nil
}

func getCycleNode(head *util.Node) *util.Node {
	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
