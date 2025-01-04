package validate_binary_search_tree

import (
	"github.com/timur-makarov/leetcode/Implementations/tree"
	"math"
)

type TreeNode = tree.TreeNode

func RunSolution(root *TreeNode) bool {
	return traverse_recursive(root, math.MinInt, math.MaxInt)
}

func traverse_recursive(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	return traverse_recursive(node.Left, minVal, node.Val) &&
		traverse_recursive(node.Right, node.Val, maxVal)
}
