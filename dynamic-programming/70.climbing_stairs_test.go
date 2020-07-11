package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "empty input",
			input:  0,
			expect: 1,
		},
		{
			name:   "2 stairs",
			input:  2,
			expect: 2,
		}, {
			name:   "3 stairs",
			input:  3,
			expect: 3,
		},
	}

	for _, c := range testTable {
		assert.Equal(t, c.expect, climbStairs(c.input))
	}
}
