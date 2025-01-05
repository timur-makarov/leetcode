package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/letter_case_permutation"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	s := "a1b2"
	answer := []string{"a1b2", "A1b2", "a1B2", "A1B2"}
	result := letter_case_permutation.RunSolution(s)
	assert.Equal(t, answer, result)
}
