package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	testTable := []struct {
		name     string
		s        string
		wordDict []string
		expect   bool
	}{
		{
			name:     "happy path",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expect:   true,
		}, {
			name:     "happy path2",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expect:   true,
		}, {
			name:     "happy path3",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expect:   true,
		}, {
			name:     "false",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expect:   false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expect, wordBreak(c.s, c.wordDict))
		})
	}
}
