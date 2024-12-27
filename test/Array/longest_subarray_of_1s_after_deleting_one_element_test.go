package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/longest_subarray_of_1s_after_deleting_one_element"
	"testing"
)

func TestLongestSubarrayof1sAfterDeletingOneElement(t *testing.T) {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	answer := 5
	result := longest_subarray_of_1s_after_deleting_one_element.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{1, 1, 1}
	answer = 2
	result = longest_subarray_of_1s_after_deleting_one_element.RunSolution(nums)
	assert.Equal(t, answer, result)
}
