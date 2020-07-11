package util

import (
	"fmt"
	"os"
)

type Node struct {
	Val  int
	Next *Node
}

const maxLoop = 30

func GenLinkedList(arr []int) *Node {
	return arrayToLinkedList(arr)
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

func GenIntersectionLinkedList(arr1, arr2, common []int) (*Node, *Node) {
	if len(common) == 0 {
		return arrayToLinkedList(arr1), arrayToLinkedList(arr2)
	}

	list1 := arrayToLinkedList(arr1)
	list2 := arrayToLinkedList(arr2)
	intersection := arrayToLinkedList(common)

	last1, last2 := getLastNode(list1), getLastNode(list2)
	last1.Next = intersection
	last2.Next = intersection

	return list1, list2
}

func GetLength(head *Node) int {
	current := head
	length := 0
	for current != nil {
		if length > maxLoop {
			fmt.Printf("\n%s\n", "max loop exceed")
			os.Exit(1)
		}
		length++
		current = current.Next
	}
	return length
}

func PrintLinkedList(node *Node) {
	format := "%v -> "
	current := node
	count := 0
	for current != nil {
		if count > maxLoop {
			fmt.Printf("\n%s\n", "loop time exceed")
			os.Exit(1)
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

func arrayToLinkedList(arr []int) *Node {
	head := &Node{}
	current := head
	for _, v := range arr {
		current.Next = &Node{v, nil}
		current = current.Next
	}
	return head.Next
}

func getLastNode(head *Node) *Node {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	return current
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
