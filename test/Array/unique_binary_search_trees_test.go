package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/unique_binary_search_trees"
	"testing"
)

func TestUniqueBinarySearchTrees(t *testing.T) {
	n := 4
	answer := 14
	result := unique_binary_search_trees.RunSolution(n)
	assert.Equal(t, answer, result)
}
