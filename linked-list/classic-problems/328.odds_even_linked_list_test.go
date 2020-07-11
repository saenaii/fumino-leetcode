package classic_problems

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenList(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 4, 5},
			expect: []int{1, 3, 5, 2, 4},
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
			head := util.GenLinkedList(c.input)
			expect := oddEvenList(head)
			assert.Equal(t, c.expect, util.LinkedListToArray(expect))
		})
	}
}
