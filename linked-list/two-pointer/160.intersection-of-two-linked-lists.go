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
func getIntersectionNode(headA, headB *util.Node) *util.Node {
	if headA == nil || headB == nil {
		return nil
	}

	lastA := getLastNode(headA)
	lastA.Next = headB
	defer func() {
		lastA.Next = nil
	}()

	if cycleNode := _getCycleNode(headA); cycleNode != nil {
		slow, fast := headA, cycleNode.Next
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

func getLastNode(head *util.Node) *util.Node {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

func _getCycleNode(head *util.Node) *util.Node {
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
