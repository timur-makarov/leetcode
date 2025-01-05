package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/subsets"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	answer := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
		{4},
		{1, 4},
		{2, 4},
		{1, 2, 4},
		{3, 4},
		{1, 3, 4},
		{2, 3, 4},
		{1, 2, 3, 4},
	}
	result := subsets.RunSolution(nums)
	assert.Equal(t, answer, result)
}
