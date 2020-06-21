package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	testTable := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name: "happy path",
			input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expect: 7,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, minPathSum(c.input))
		})
	}
}
