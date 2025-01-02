package Tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Tree/all_nodes_distance_k_in_binary_tree"
	"testing"
)

func TestAllNodesDistanceKinBinaryTree(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(25)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	answer := []int{20}
	result := all_nodes_distance_k_in_binary_tree.RunSolution(tr.Root, tr.Root.Left, 2)
	assert.Equal(t, answer, result)
}
