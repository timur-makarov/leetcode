package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/non_overlapping_intervals"
	"testing"
)

func TestNonOverlappingIntervals(t *testing.T) {
	intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	answer := 2
	result := non_overlapping_intervals.RunSolution(intervals)
	assert.Equal(t, answer, result)
}
