package util

import "fmt"

type Node struct {
	Val int
	Next *Node
}

const maxLoop = 30

func GenLinkedList(arr []int) *Node {
	head := &Node{}
	current := head
	for _, v := range arr {
		current.Next = &Node{v, nil}
		current = current.Next
	}
	return head.Next
}

func LinkedListToArray(node *Node) []int {
	arr := make([]int, 0)
	current := node
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

func GenLinkedListWithCycle(arr []int, cycle int) *Node {
	head := &Node{}
	current := head
	target := &Node{}
	for i, v := range arr {
		node := &Node{v, nil}
		if i == cycle {
			target = node
		}
		current.Next = node
		current = current.Next
	}
	current.Next = target
	return head.Next
}

func PrintLinkedList(node *Node) {
	format := "%v -> "
	current := node
	count := 0
	for current != nil {
		if count > maxLoop {
			break
		}
		if current.Next == nil {
			format = "%v"
		}
		fmt.Printf(format, current.Val)
		current = current.Next
		count++
	}
	fmt.Println()
}

func getValByIndex(head *Node, index int) int {
	if index == -1 {
		return -1
	}
	current := head
	for i := 0; i < index; i++ {
		if current == nil {
			return -1
		}
		current = current.Next
	}
	return current.Val
}

