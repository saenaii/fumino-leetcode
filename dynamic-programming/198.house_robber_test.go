package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{1, 2, 3, 1},
			expect: 4,
		}, {
			name:   "happy path2",
			input:  []int{2, 7, 9, 3, 1},
			expect: 12,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: 0,
		}, {
			name:   "one element",
			input:  []int{1},
			expect: 1,
		}, {
			name:   "2 elements",
			input:  []int{1, 2},
			expect: 2,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, rob(c.input))
			assert.Equal(t, c.expect, robSolution2(c.input))
		})
	}
}
