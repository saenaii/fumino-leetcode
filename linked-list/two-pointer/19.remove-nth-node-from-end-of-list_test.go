package two_pointer

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		nth    int
		expect []int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 4, 5},
			nth:    2,
			expect: []int{1, 2, 3, 5},
		}, {
			name:   "empty input",
			input:  []int{},
			nth:    1,
			expect: []int{},
		}, {
			name:   "remove first",
			input:  []int{1, 2, 3},
			nth:    3,
			expect: []int{2, 3},
		}, {
			name:   "remove last",
			input:  []int{1, 2, 3, 4},
			nth:    1,
			expect: []int{1, 2, 3},
		}, {
			name:   "one element",
			input:  []int{1},
			nth:    1,
			expect: []int{},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			head := util.GenLinkedList(c.input)
			head = removeNthFromEnd(head, c.nth)
			assert.Equal(t, c.expect, util.LinkedListToArray(head))
		})
	}
}
