package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/find_peak_element"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 12, 15}
	answer := -1
	result := find_peak_element.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	answer = 3
	result = find_peak_element.RunSolution(nums)
	assert.Equal(t, answer, result)
}
