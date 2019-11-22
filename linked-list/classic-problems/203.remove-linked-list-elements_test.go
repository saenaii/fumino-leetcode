package classic_problems

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		val    int
		expect []int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 6, 3, 4, 5, 6},
			val:    6,
			expect: []int{1, 2, 3, 4, 5},
		}, {
			name:   "one element",
			input:  []int{1},
			val:    1,
			expect: []int{},
		}, {
			name:   "all same element",
			input:  []int{1, 1, 1, 1, 1},
			val:    1,
			expect: []int{},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			linkedList := util.GenLinkedList(c.input)
			head := removeElements(linkedList, c.val)
			assert.Equal(t, c.expect, util.LinkedListToArray(head))
		})
	}
}
