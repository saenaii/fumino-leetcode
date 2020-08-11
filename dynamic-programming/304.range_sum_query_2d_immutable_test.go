package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRegion(t *testing.T) {
	testTable := []struct{
		name string
		matrix [][]int
		params []struct{
			row1, col1, row2, col2, expect int
		}
	} {
		{
			name: "happy path",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			params: []struct{ row1, col1, row2, col2, expect int }{
				{row1: 2, col1: 1, row2: 4, col2: 3,expect: 8},
				{row1: 1, col1: 1, row2: 2, col2: 2, expect: 11},
				{row1: 1, col1: 2, row2: 2, col2: 4, expect: 12},
			},
		},
	}

	for _, c := range testTable{
		t.Run(c.name, func(t *testing.T) {
			obj := ConstructorNumMatrix(c.matrix)
			for _, item := range c.params{
				assert.Equal(t, item.expect, obj.SumRegion(item.row1, item.col1, item.row2, item.col2))
			}

			obj2 := ConstructorNumMatrix2(c.matrix)
			for _, item := range c.params{
				assert.Equal(t, item.expect, obj2.SumRegion(item.row1, item.col1, item.row2, item.col2))
			}
		})
	}
}
