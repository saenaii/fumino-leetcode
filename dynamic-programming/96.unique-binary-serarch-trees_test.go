package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmTrees(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "happy path",
			input:  3,
			expect: 5,
		}, {
			name:   "happy path2",
			input:  6,
			expect: 132,
		},
	}

	for _, c := range testTable {
		assert.Equal(t, c.expect, numTrees(c.input))
		assert.Equal(t, c.expect, numTreesSolution2(c.input))
	}
}
