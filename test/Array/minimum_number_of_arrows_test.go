package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/minimum_number_of_arrows"
	"testing"
)

func TestMinimumNumberofArrows(t *testing.T) {
	intervals := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	answer := 2
	result := minimum_number_of_arrows.RunSolution(intervals)
	assert.Equal(t, answer, result)

	intervals = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	answer = 2
	result = minimum_number_of_arrows.RunSolution(intervals)
	assert.Equal(t, answer, result)

	intervals = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	answer = 4
	result = minimum_number_of_arrows.RunSolution(intervals)
	assert.Equal(t, answer, result)
}
