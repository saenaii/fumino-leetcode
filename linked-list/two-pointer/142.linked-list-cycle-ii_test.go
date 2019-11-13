package two_pointer

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	testTable := []struct{
		name string
		input []int
		cycle int
		expect int
	} {
		{
			name: "happy path",
			input: []int{1,2,3,4,5},
			cycle: 1,
			expect: 2,
		}, {
			name: "empty input",
			input: []int{},
			cycle:-1,
			expect:0,
		}, {
			name: "no cycle",
			input: []int{1,2,3,4,5},
			cycle:-1,
			expect:-1,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			linkedList := util.GenLinkedListWithCycle(c.input, c.cycle)
			if node := detectCycle(linkedList); node != nil {
				assert.Equal(t, c.expect, node.Val)
			}
		})
	}
}
