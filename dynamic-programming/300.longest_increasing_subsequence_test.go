package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		}, {
			name:   "happy path2",
			input:  []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expect: 6,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, lengthOfLIS(c.input))
		})
	}
}
