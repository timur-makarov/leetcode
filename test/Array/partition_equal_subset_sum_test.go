package Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/partition_equal_subset_sum"
)

func TestPartitionEqualSubsetSum(t *testing.T) {
	nums := []int{1, 5, 5, 11}
	answer := true
	result := partition_equal_subset_sum.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, 3, 5}
	answer = false
	result = partition_equal_subset_sum.RunSolution(nums)
	assert.Equal(t, answer, result)
}
