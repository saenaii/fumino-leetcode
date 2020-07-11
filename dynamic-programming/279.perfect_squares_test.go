package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "happy path",
			input:  12,
			expect: 3,
		}, {
			name:   "happy path2",
			input:  13,
			expect: 2,
		}, {
			name:   "happy path3",
			input:  1,
			expect: 1,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, numSquares(c.input))
			assert.Equal(t, c.expect, numSquaresSolution2(c.input))
		})
	}

}
