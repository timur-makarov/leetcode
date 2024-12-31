package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/shortest_unsorted_continuous_subarray"
	"testing"
)

func TestShortestUnsortedContinuousSubarray(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	answer := 5
	result := shortest_unsorted_continuous_subarray.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, 3, 4}
	answer = 0
	result = shortest_unsorted_continuous_subarray.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{4, 3, 2, 1}
	answer = 4
	result = shortest_unsorted_continuous_subarray.RunSolution(nums)
	assert.Equal(t, answer, result)
}
