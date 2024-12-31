package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/permutation_in_string"
	"testing"
)

func TestPermutationInString(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	result := permutation_in_string.RunSolution(s1, s2)
	assert.Equal(t, true, result)

	s1 = "ab"
	s2 = "eidboaoo"
	result = permutation_in_string.RunSolution(s1, s2)
	assert.Equal(t, false, result)
}
