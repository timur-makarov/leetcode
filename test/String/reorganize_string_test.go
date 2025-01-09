package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/reorganize_string"
	"testing"
)

func TestReorganizeString(t *testing.T) {
	s := "aab"
	answer := "aba"
	result := reorganize_string.RunSolution(s)
	assert.Equal(t, answer, result)

	s = "aaab"
	answer = ""
	result = reorganize_string.RunSolution(s)
	assert.Equal(t, answer, result)

	s = "aabbb"
	answer = "babab"
	result = reorganize_string.RunSolution(s)
	assert.Equal(t, answer, result)
}
