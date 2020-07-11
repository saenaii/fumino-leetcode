package classic_problems

import (
	"fumino-leetcode/linked-list/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeLinkedList(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect bool
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 2, 1},
			expect: true,
		}, {
			name:   "odds elements",
			input:  []int{1, 2, 1},
			expect: true,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: true,
		}, {
			name:   "not palindrome",
			input:  []int{1, 2},
			expect: false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			head := util.GenLinkedList(c.input)
			result := isPalindrome(head)
			assert.Equal(t, c.expect, result)
		})
	}
}
