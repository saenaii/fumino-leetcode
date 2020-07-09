package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "happy path",
			input:  10,
			expect: 12,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, nthUglyNumber(c.input))
			assert.Equal(t, c.expect, nthUglyNumberSolution2(c.input))
		})
	}
}
