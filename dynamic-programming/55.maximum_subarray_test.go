package dynamic_programming

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		}, {
			name:   "negative input",
			input:  []int{-3, -4, -1, -5, -100},
			expect: -1,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: math.MinInt32,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, maxSubArray(c.input))
		})
	}
}
