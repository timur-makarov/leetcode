package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/permutations"
	"testing"
)

func TestPermutations(t *testing.T) {
	nums := []int{1, 2, 3}
	answer := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	result := permutations.RunSolution(nums)
	assert.Equal(t, answer, result)
}
