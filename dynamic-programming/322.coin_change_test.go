package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	testTable := []struct{
		name string
		input []int
		amount int
		expect int
	} {
		{
			name: "happy path",
			input: []int{1, 2, 5},
			amount: 11,
			expect: 3,
		},
		{
			name:"happy path2",
			input: []int{186,419,83,408},
			amount: 6249,
			expect: 20,
		},
		{
			name: "can not make up",
			input: []int{2},
			amount: 3,
			expect: -1,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, coinChange(c.input, c.amount))
		})
	}
}