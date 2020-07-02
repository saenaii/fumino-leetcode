package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob2(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{2, 3, 2},
			expect: 3,
		}, {
			name:   "happy path2",
			input:  []int{1, 2, 3, 1},
			expect: 4,
		}, {
			name:   "happy path3",
			input:  []int{1, 2, 1, 1},
			expect: 3,
		}, {
			name:   "empty input",
			input:  []int{},
			expect: 0,
		}, {
			name:   "1 element",
			input:  []int{2},
			expect: 2,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, rob2(c.input))
		})
	}
}
