package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/maximum_number_of_distinct_elements_after_operations"
	"testing"
)

func TestMaximumNumberOfDistinctElementsAfterOperations(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 4}
	k := 2
	answer := 6
	result := maximum_number_of_distinct_elements_after_operations.RunSolution(nums, k)
	assert.Equal(t, answer, result)

	nums = []int{4, 4, 4, 4}
	k = 1
	answer = 3
	result = maximum_number_of_distinct_elements_after_operations.RunSolution(nums, k)
	assert.Equal(t, answer, result)

	nums = []int{8, 8, 10, 9, 9}
	k = 1
	answer = 5
	result = maximum_number_of_distinct_elements_after_operations.RunSolution(nums, k)
	assert.Equal(t, answer, result)
}
