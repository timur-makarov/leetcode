package Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/knapsack"
)

func TestKnapsack(t *testing.T) {
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	capacity := 7
	answer := 22
	result := knapsack.RunSolution(profits, weights, capacity)
	assert.Equal(t, answer, result)
}
