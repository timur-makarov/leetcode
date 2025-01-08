package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/k_closest_points_to_origin"
	"testing"
)

func TestKClosestPointsToOrigin(t *testing.T) {
	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k := 2
	answer := [][]int{{3, 3}, {-2, 4}}
	result := k_closest_points_to_origin.RunSolution(points, k)
	assert.ElementsMatch(t, answer, result)

	points = [][]int{{1, 1}, {5, -3}, {0, 2}}
	k = 1
	answer = [][]int{{1, 1}}
	result = k_closest_points_to_origin.RunSolution(points, k)
	assert.ElementsMatch(t, answer, result)
}
