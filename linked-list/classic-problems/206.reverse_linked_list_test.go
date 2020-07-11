package classic_problems

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 4, 5},
			expect: []int{5, 4, 3, 2, 1},
		}, {
			name:   "empty input",
			input:  []int{},
			expect: []int{},
		}, {
			name:   "one element",
			input:  []int{1},
			expect: []int{1},
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			head := reverseList(util.GenLinkedList(c.input))
			assert.Equal(t, c.expect, util.LinkedListToArray(head))
		})
	}
}
