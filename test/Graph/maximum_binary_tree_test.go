package Graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Graph/maximum_binary_tree"
)

func TestMaximumBinaryTree(t *testing.T) {
	nums := []int{9, 20, 10, 1, 3, 4, 15, 10, 25}
	answer := 25
	root := maximum_binary_tree.RunSolution(nums)
	assert.Equal(t, answer, root.Val)
}
