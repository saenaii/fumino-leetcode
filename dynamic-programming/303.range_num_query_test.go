package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	testTable := []struct {
		name string
		i, j int
		expect int
	} {
		{
			name: "happy path",
			i: 0,
			j:2,
			expect: 1,
		}, {
			name: "happy path2",
			i:2,
			j:5,
			expect: -1,
		}, {
			name:"happy path3",
			i:0,
			j:5,
			expect: -3,
		},
	}

	obj := Constructor(nums)
	for _, c :=range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, obj.SumRange(c.i, c.j))
		})
	}
}