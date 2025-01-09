package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/maximum_frequency_stack"
	"testing"
)

func TestMaximumFrequencyStack(t *testing.T) {
	answer := []int{5, 7, 5, 4}
	result := maximum_frequency_stack.RunSolution()
	assert.Equal(t, answer, result)
}
