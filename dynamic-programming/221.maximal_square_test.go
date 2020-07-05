package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {
	testTable := []struct {
		name   string
		input  [][]byte
		expect int
	}{
		{
			name: "happy path",
			input: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expect: 4,
		}, {
			name: "happy path2",
			input: [][]byte{
				{'1', '0', '1', '0'},
				{'1', '0', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
			},
			expect: 4,
		}, {
			name: "1 element",
			input: [][]byte{
				{'1'},
			},
			expect: 1,
		}, {
			name:   "empty input",
			input:  [][]byte{},
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, maximalSquare(c.input))
			assert.Equal(t, c.expect, maximalSquareSolution2(c.input))
		})
	}
}
