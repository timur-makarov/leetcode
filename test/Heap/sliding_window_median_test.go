package Heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/sliding_window_median"
	"testing"
)

func TestSlidingWindowMedian(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 3
	answer := []float64{2, 3, 4, 5}
	result := sliding_window_median.RunSolution(nums, k)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, 3}
	k = 2
	answer = []float64{1.5, 2.5}
	result = sliding_window_median.RunSolution(nums, k)
	assert.Equal(t, answer, result)

	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = 3
	answer = []float64{1, -1, -1, 3, 5, 6}
	result = sliding_window_median.RunSolution(nums, k)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k = 3
	answer = []float64{2, 3, 3, 3, 2, 3, 2}
	result = sliding_window_median.RunSolution(nums, k)
	assert.Equal(t, answer, result)
}
