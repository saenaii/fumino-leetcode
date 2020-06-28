package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "happy path",
			input:  []int{2, 3, -2, 4},
			expect: 6,
		}, {
			name:   "happy path2",
			input:  []int{-2, 0, -1},
			expect: 0,
		}, {
			name:   "happy path3",
			input:  []int{-2, 3, -4},
			expect: 24,
		}, {
			name:   "happy path4",
			input:  []int{0, 2},
			expect: 2,
		}, {
			name:   "happy path5",
			input:  []int{2, -5, -2, -4, 3},
			expect: 24,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, maxProduct(c.input))
		})
	}
}
