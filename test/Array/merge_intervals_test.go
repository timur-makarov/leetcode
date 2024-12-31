package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/merge_intrevals"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	answer := [][]int{{1, 6}, {8, 10}, {15, 18}}
	result := merge_intrevals.RunSolution(nums)
	assert.Equal(t, answer, result)
}
