package Tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Implementations/tree"
	"github.com/timur-makarov/leetcode/Tree/binary_tree_level_order_traversal"
	"testing"
)

type Tree = tree.Tree
type TreeNode = tree.TreeNode

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	var tr Tree

	_ = tr.Insert(9)
	_ = tr.Insert(3)
	_ = tr.Insert(20)
	_ = tr.Insert(15)
	_ = tr.Insert(25)

	answer := [][]int{{15, 25}, {3, 20}, {9}}
	result := binary_tree_level_order_traversal.RunSolution(tr.Root)
	assert.Equal(t, answer, result)
}
