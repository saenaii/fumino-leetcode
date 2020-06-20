package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testTable := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name: "happy path",
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expect: 2,
		}, {
			name: "happy path2",
			input: [][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			},
			expect: 0,
		}, {
			name: "one line",
			input: [][]int{
				{1, 0},
			},
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, uniquePathsWithObstacles(c.input))
		})
	}
}
