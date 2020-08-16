package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit309(t *testing.T) {
	testTable := []struct{
		name string
		input []int
		expect int
	}{
		{
			name: "happy path",
			input: []int{1,2,3,0,2},
			expect: 3,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, maxProfit309(c.input))
		})
	}
}
