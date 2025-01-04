package Heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/find_median_from_data_stream"
	"testing"
)

func TestFindMedianfromDataStream(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	answer := 3.5
	result := find_median_from_data_stream.RunSolution(nums)
	assert.Equal(t, answer, result)

	nums = []int{1, 2, 3}
	answer = float64(2)
	result = find_median_from_data_stream.RunSolution(nums)
	assert.Equal(t, answer, result)
}
