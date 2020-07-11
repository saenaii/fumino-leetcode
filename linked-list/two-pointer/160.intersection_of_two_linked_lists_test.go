package two_pointer

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	testTable := []struct {
		name            string
		list1           []int
		list2           []int
		common          []int
		expect          int
		hasIntersection bool
	}{
		{
			name:            "happy path",
			list1:           []int{4, 1},
			list2:           []int{5, 0, 1},
			common:          []int{8, 4, 5},
			expect:          8,
			hasIntersection: true,
		}, {
			name:            "no intersection",
			list1:           []int{1, 3, 4},
			list2:           []int{3, 4, 5},
			common:          []int{},
			expect:          -1,
			hasIntersection: false,
		}, {
			name:            "empty input",
			list1:           []int{},
			list2:           []int{1, 2, 3},
			common:          []int{},
			expect:          -1,
			hasIntersection: false,
		}, {
			name:            "one element in list",
			list1:           []int{0, 9, 1},
			list2:           []int{3},
			common:          []int{2, 4},
			expect:          2,
			hasIntersection: true,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			headA, headB := util.GenIntersectionLinkedList(c.list1, c.list2, c.common)
			node := getIntersectionNode(headA, headB)
			if c.hasIntersection {
				assert.Equal(t, c.expect, node.Val)
				assert.Equal(t, len(c.list1)+len(c.common), util.GetLength(headA))
				assert.Equal(t, len(c.list2)+len(c.common), util.GetLength(headB))
			} else {
				assert.Nil(t, node)
			}
			assert.Equal(t, len(c.list1)+len(c.common), util.GetLength(headA))
			assert.Equal(t, len(c.list2)+len(c.common), util.GetLength(headB))
		})
	}
}
