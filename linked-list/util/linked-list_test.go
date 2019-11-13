package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenLinkedList(t *testing.T) {
	testTable := []struct{
		name string
		input []int
	} {
		{
			name: "happy path",
			input: []int{1,2,3,4,5},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			node := GenLinkedList(c.input)
			PrintLinkedList(node)
		})
	}
}

func TestGenLinkedListWithCycle(t *testing.T) {
	testTable := []struct {
		name string
		input []int
		cycle int
		target int
	} {
		{
			name:"happy path",
			input: []int{1, 2, 3, 4, 5},
			cycle:2,
			target: 3,
		}, {
			name: "no cycle",
			input: []int{1,3,4,6},
			cycle:-1,
			target:-1,
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
	testTable := []struct{
		name string
		input []int
	} {
		{
			name: "happy path",
			input: []int{1,2,3,4,5},
		}, {
			name: "empty input",
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