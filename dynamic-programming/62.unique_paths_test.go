package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	testTable := []struct {
		name   string
		m, n   int
		expect int
	}{
		{
			name:   "happy path",
			m:      3,
			n:      2,
			expect: 3,
		}, {
			name:   "happy path2",
			m:      7,
			n:      3,
			expect: 28,
		},
	}

	for _, c := range testTable {
		assert.Equal(t, c.expect, uniquePaths(c.m, c.n))
		assert.Equal(t, c.expect, uniquePathsSolution2(c.m, c.n))
	}
}
