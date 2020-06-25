package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	testTable := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name: "happy path",
			input: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expect: 11,
		}, {
			name: "happy path2",
			input: [][]int{
				{1},
			},
			expect: 1,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, minimumTotal(c.input))
			assert.Equal(t, c.expect, minimumTotalSolution2(c.input))
		})
	}
}
