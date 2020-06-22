package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDecodings(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "happy path",
			input:  "226",
			expect: 3,
		}, {
			name:   "happy path2",
			input:  "12",
			expect: 2,
		}, {
			name:   "edge case",
			input:  "0",
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, numDecodings(c.input))
		})
	}
}
