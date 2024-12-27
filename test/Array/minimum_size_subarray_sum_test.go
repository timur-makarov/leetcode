package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/minimum_size_subarray_sum"
	"testing"
)

func TestMinimumSizeSubarraySum(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	result := minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 2, result)

	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	result = minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 0, result)

	target = 11
	nums = []int{1, 2, 3, 4, 5}
	result = minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 3, result)

	target = 15
	nums = []int{1, 2, 3, 4, 5}
	result = minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 5, result)

	target = 1000000000
	nums = make([]int, 100000)
	for i := 0; i < 100000; i++ {
		nums[i] = 10000
	}
	result = minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 100000, result)

	target = 7
	nums = []int{8}
	result = minimum_size_subarray_sum.RunSolution(target, nums)
	assert.Equal(t, 1, result)
}
