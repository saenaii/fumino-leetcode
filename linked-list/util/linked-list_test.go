package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenLinkedList(t *testing.T) {
	testTable := []struct {
		name  string
		input []int
	}{
		{
			name:  "happy path",
			input: []int{1, 2, 3, 4, 5},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			node := GenLinkedList(c.input)
			assert.Equal(t, len(c.input), GetLength(node))
		})
	}
}

func TestGenLinkedListWithCycle(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		cycle  int
		target int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 4, 5},
			cycle:  2,
			target: 3,
		}, {
			name:   "no cycle",
			input:  []int{1, 3, 4, 6},
			cycle:  -1,
			target: -1,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			node := GenLinkedListWithCycle(c.input, c.cycle)
			assert.Equal(t, c.target, getValByIndex(node, c.cycle))
		})
	}
}

func TestLinkedListToArray(t *testing.T) {
	testTable := []struct {
		name  string
		input []int
	}{
		{
			name:  "happy path",
			input: []int{1, 2, 3, 4, 5},
		}, {
			name:  "empty input",
			input: []int{},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			node := GenLinkedList(c.input)
			assert.Equal(t, c.input, LinkedListToArray(node))
		})
	}
}

func TestGenIntersectionLinkedList(t *testing.T) {
	testTable := []struct {
		name            string
		list1           []int
		list2           []int
		common          []int
		hasIntersection bool
	}{
		{
			name:            "happy path",
			list1:           []int{4, 1},
			list2:           []int{5, 0, 1},
			common:          []int{8, 4, 5},
			hasIntersection: true,
		}, {
			name:            "no intersection",
			list1:           []int{1, 2, 3},
			list2:           []int{2, 3, 4},
			common:          []int{},
			hasIntersection: false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			head1, head2 := GenIntersectionLinkedList(c.list1, c.list2, c.common)
			if c.hasIntersection {
				assert.Equal(t, getLastNode(head1), getLastNode(head2))
			} else {
				assert.NotEqual(t, getLastNode(head1), getLastNode(head2))
			}
		})
	}
}

func TestGetLength(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 4},
			expect: 4,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			head := arrayToLinkedList(c.input)
			assert.Equal(t, c.expect, GetLength(head))
		})
	}
}
