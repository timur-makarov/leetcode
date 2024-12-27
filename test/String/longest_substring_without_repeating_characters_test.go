package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/longest_substring_without_repeating_characters"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	s := "abcabcbb"
	answer := 3
	result := longest_substring_without_repeating_characters.RunSolution(s)
	assert.Equal(t, answer, result)

	s = "bbbbb"
	answer = 1
	result = longest_substring_without_repeating_characters.RunSolution(s)
	assert.Equal(t, answer, result)

	s = "pwwkew"
	answer = 3
	result = longest_substring_without_repeating_characters.RunSolution(s)
	assert.Equal(t, answer, result)

	s = "au"
	answer = 2
	result = longest_substring_without_repeating_characters.RunSolution(s)
	assert.Equal(t, answer, result)
}
