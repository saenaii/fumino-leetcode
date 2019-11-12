package singly_linked_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList_Get(t *testing.T) {
	linkedList := Constructor()
	input := []int{1,3,5,4}
	for _, v := range input {
		linkedList.AddAtTail(v)
	}

	testTable := []struct {
		name string
		index int
		expect int
	} {
		{
			name: "first element",
			index: 0,
			expect:1,
		}, {
			name: "out of range",
			index: 10,
			expect: -1,
		}, {
			name: "last element",
			index: 3,
			expect: 4,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, linkedList.Get(c.index))
		})
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	linkedList := Constructor()
	input := []int{2, 3, 5}
	for _, v := range input {
		linkedList.AddAtHead(v)
		assert.Equal(t, linkedList.head.val, v)
		assert.Nil(t, linkedList.last.next)
	}

	assert.NotNil(t, linkedList.head.next)
	assert.Equal(t, len(input), linkedList.length)
	assert.Equal(t, input[0], linkedList.last.val)
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	linkedList := Constructor()
	input := []int{3,2,5,6,7}
	for _, v := range input {
		linkedList.AddAtTail(v)
		assert.Equal(t, linkedList.last.val, v)
		assert.Nil(t, linkedList.last.next)
	}

	assert.Equal(t, len(input), linkedList.length)
	assert.Equal(t, input[0], linkedList.head.val)
	assert.NotNil(t, linkedList.head.next)
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	linkedList := Constructor()
	testTable := []struct{
		name string
		index int
		value int
		expectValue int
		expectFirst int
		expectLast int
		expectLength int
	} {
		{
			name: "first insert",
			index: 0,
			value: 3,
			expectValue: 3,
			expectFirst: 3,
			expectLast: 3,
			expectLength: 1,
		}, {
			name: "out of range",
			index: 10,
			value: 3,
			expectValue: -1,
			expectFirst:3,
			expectLast: 3,
			expectLength: 1,
		}, {
			name: "add at tail",
			index: 1,
			value: 7,
			expectValue: 7,
			expectFirst:3,
			expectLast:7,
			expectLength:2,
		}, {
			name: "add at middle",
			index: 1,
			value: 4,
			expectValue: 4,
			expectFirst:3,
			expectLast:7,
			expectLength:3,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			linkedList.AddAtIndex(c.index, c.value)
			assert.Equal(t, c.expectFirst, linkedList.head.val)
			assert.Equal(t, c.expectLast, linkedList.last.val)
			assert.Equal(t, c.expectLength, linkedList.length)
			assert.Equal(t, c.expectValue, linkedList.Get(c.index))
		})
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	linkedList := Constructor()
	input := []int{3,2,5,6,7}
	for _, v := range input {
		linkedList.AddAtTail(v)
	}

	testTable := []struct{
		name string
		index int
		expectFirst int
		expectLast int
		expectLength int
	} {
		{
			name: "delete first",
			index: 0,
			expectFirst: 2,
			expectLast: 7,
			expectLength: 4,
		}, {
			name: "out of range",
			index: 100,
			expectFirst: 2,
			expectLast: 7,
			expectLength: 4,
		}, {
			name: "delete last",
			index: 3,
			expectFirst: 2,
			expectLast:6,
			expectLength: 3,
		}, {
			name: "delete middle",
			index: 1,
			expectFirst: 2,
			expectLast: 6,
			expectLength: 2,
		},
	}


	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			linkedList.DeleteAtIndex(c.index)
			assert.Equal(t, c.expectFirst, linkedList.head.val)
			assert.Equal(t, c.expectLast, linkedList.last.val)
			assert.Equal(t, c.expectLength, linkedList.length)
		})
	}
}

type LinkedList struct {
	head *Node
}

func TestConstructor(t *testing.T) {
	l := new(LinkedList)
l = AddAtHead(3, l)
	l = AddAtHead(2, l)
	l = AddAtHead(5, l)
	printList(l)
}

func printList(l *LinkedList) {
	current := l.head
	for current != nil {
		fmt.Printf("%v->", current.val)
		current = current.next
	}
	fmt.Println()
}

func AddAtHead(val int, h *LinkedList) *LinkedList {
	node := &Node{val, nil}
	if h.head == nil {
		h.head = node
		return h
	}

	current := h.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	return h
}