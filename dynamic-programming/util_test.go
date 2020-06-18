package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 1, max(1, 0))
	assert.Equal(t, 2, max(1, 2))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 100))
	assert.Equal(t, 1, min(100, 1))
}
