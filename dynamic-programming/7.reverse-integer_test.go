package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testTable := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "happy path",
			input:  123,
			expect: 321,
		}, {
			name:   "negative input",
			input:  -123,
			expect: -321,
		}, {
			name:   "happy path2",
			input:  120,
			expect: 21,
		}, {
			name:   "out of scope",
			input:  1233456344,
			expect: 0,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, reverse(c.input))
		})
	}
}
