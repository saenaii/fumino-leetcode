package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "empty input",
			input:  []int{},
			expect: 0,
		}, {
			name:   "happy path",
			input:  []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
	}

	for _, c := range testTable {
		assert.Equal(t, c.expect, maxProfit(c.input))
	}
}
