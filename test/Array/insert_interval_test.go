package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/insert_interval"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	newInterval := []int{12, 14}
	answer := [][]int{{1, 3}, {2, 6}, {8, 10}, {12, 14}, {15, 18}}
	result := insert_interval.RunSolution(intervals, newInterval)
	assert.Equal(t, answer, result)
}
