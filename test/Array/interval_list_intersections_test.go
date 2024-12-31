package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/interval_list_intersections"
	"testing"
)

func TestIntervalListIntersections(t *testing.T) {
	firstIntervals := [][]int{{2, 6}, {7, 12}}
	secondIntervals := [][]int{{1, 3}, {8, 10}, {15, 18}}
	answer := [][]int{{2, 3}, {8, 10}}
	result := interval_list_intersections.RunSolution(firstIntervals, secondIntervals)
	assert.Equal(t, answer, result)
}
