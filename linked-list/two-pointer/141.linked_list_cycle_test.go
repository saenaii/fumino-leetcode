package two_pointer

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	testTable := []struct{
		name string
		input []int
		cycle int
		hasCycle bool
	}{
		{
			name: "no cycle",
			input:[]int{1,2,3,4},
			cycle:-1,
			hasCycle:false,
		}, {
			name: "cycle at first",
			input:[]int{1,3,4,5,6},
			cycle:0,
			hasCycle:true,
		}, {
			name: "empty input",
			input: []int{},
			cycle:-1,
			hasCycle:false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			linkedList := util.GenLinkedListWithCycle(c.input, c.cycle)
			assert.Equal(t, c.hasCycle, hasCycle(linkedList))
		})
	}
}
